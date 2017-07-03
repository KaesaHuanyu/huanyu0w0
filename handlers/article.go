package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"huanyu0w0/model"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func (h *Handler) CreateArticleGet(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Url string
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path=article/create")
	}
	return c.Render(http.StatusOK, "createarticle", data)
}

func (h *Handler) CreateArticle(c echo.Context) (err error) {
	//获得当前登录账户的ID
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}

	a := &model.Article{
		ID:     bson.NewObjectId(),
		Time:   time.Now(),
		Editor: data.ID,
	}

	if err = c.Bind(a); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	if a.Reason == "" || a.Title == "" || a.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "标题、作品名或安利理由为空"}
	}

	if len(a.Title) > 60 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "标题名不能超过20个字"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB(MONGO_DB).C(ARTICLE).
		Find(bson.M{"title": a.Title}).One(&model.Article{}); err == nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "该标题已存在"}
	}

	if err = db.DB(MONGO_DB).C(ARTICLE).Insert(a); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	//更新user的articles
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$addToSet": bson.M{"articles": a.ID.Hex()}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	//记录日志
	log := &model.Log{
		ID:        bson.NewObjectId(),
		Time:      time.Now(),
		Object:    a.ID.Hex(),
		Type:      "文稿",
		User:      data.ID,
		Operation: "创建",
	}

	if err = db.DB(MONGO_DB).C(LOG).
		Insert(log); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	return c.Redirect(http.StatusFound, "/article/"+a.ID.Hex())
}

func (h *Handler) ArticleDetail(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Display *model.Display
		Content template.HTML
	}{
		Display: &model.Display{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}

	id := c.Param("id")
	data.Display.Article = &model.Article{}
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB(MONGO_DB).C(ARTICLE).
		FindId(bson.ObjectIdHex(id)).
		One(data.Display.Article); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}
	data.Content = template.HTML(blackfriday.MarkdownCommon([]byte(data.Display.Article.Reason)))
	data.Display.Article.Reason = ""

	//当前用户是否点赞
	if v, ok := data.Display.Article.UserLiked[data.ID]; ok {
		if v {
			data.Display.IsLike = true
		}
	}

	//增加点击量
	data.Display.Article.Click++
	if err = db.DB(MONGO_DB).C(ARTICLE).
		UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"click": data.Display.Article.Click}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	data.Display.Editor = &model.User{}
	data.Display.ID = data.Display.Article.ID.Hex()
	data.Display.ShowTime = data.Display.Article.GetShowTime()
	data.Display.ShowTopic = data.Display.Article.GetShowTopic()
	data.Display.Article.Click = data.Display.Article.Click / 2
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(data.Display.Article.Editor)).
		One(data.Display.Editor); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}
	data.Display.Editor.Password = ""

	//当前用户是否关注
	if v, ok := data.Display.Editor.IsFollower[data.ID]; ok {
		if v {
			data.Display.IsFollow = true
		}
	}

	data.Display.MostLikes = &model.DisplayComment{
		Comment: &model.Comment{},
		Editor:  &model.User{},
		Replyto: &model.User{},
	}
	wg := sync.WaitGroup{}
	for i, v := range data.Display.Article.Comments {
		wg.Add(1)
		data.Display.Comments = append(data.Display.Comments, &model.DisplayComment{})
		go func(i int, v string) {
			defer wg.Done()
			db := h.DB.Clone()
			defer db.Close()
			data.Display.Comments[i].Number = i + 1
			data.Display.Comments[i].Comment = &model.Comment{}
			if err := db.DB(MONGO_DB).C(COMMENT).
				FindId(bson.ObjectIdHex(v)).
				One(data.Display.Comments[i].Comment); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, ":", err)
			}
			data.Display.Comments[i].ID = data.Display.Comments[i].Comment.ID.Hex()
			data.Display.Comments[i].ShowTime = data.Display.Comments[i].Comment.GetShowTime()
			data.Display.Comments[i].ReplyNum = len(data.Display.Comments[i].Comment.Replies)

			data.Display.Comments[i].Editor = &model.User{}
			if err := db.DB(MONGO_DB).C(USER).
				FindId(bson.ObjectIdHex(data.Display.Comments[i].Comment.Editor)).
				One(data.Display.Comments[i].Editor); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, " Editor:", err)
			}
			//是否为回复
			if data.Display.Comments[i].Comment.Replyto != "" {
				master := &model.Comment{}
				if err := db.DB(MONGO_DB).C(COMMENT).
					FindId(bson.ObjectIdHex(data.Display.Comments[i].Comment.Replyto)).
					One(master); err != nil {
					log.Println("<(￣︶￣)↗[GO!]", i, " Replyto:", err)
				}
				data.Display.Comments[i].Replyto = &model.User{}
				if err := db.DB(MONGO_DB).C(USER).
					FindId(bson.ObjectIdHex(master.Editor)).
					One(data.Display.Comments[i].Replyto); err != nil {
					log.Println("<(￣︶￣)↗[GO!]", i, " Replyto:", err)
				}
				data.Display.Comments[i].Replyto.Password = ""
			}
			//是否是楼主
			if data.Display.Comments[i].Comment.Editor == data.Display.Article.Editor {
				data.Display.Comments[i].IsEditor = true
			}

			//筛选最热评论
			if data.Display.Comments[i].Comment.Like > data.Display.MostLikes.Comment.Like {
				data.Display.MostLikes = data.Display.Comments[i]
			}

			//当前用户是否点赞
			if v, ok := data.Display.Comments[i].Comment.UserLiked[data.ID]; ok {
				if v {
					data.Display.Comments[i].IsLike = true
				}
			}

			//当前用户是否关注
			if v, ok := data.Display.Comments[i].Editor.IsFollower[data.ID]; ok {
				if v {
					data.Display.Comments[i].IsFollow = true
				}
			}
		}(i, v)
	}
	wg.Wait()
	if data.Display.MostLikes.Comment.Like >= len(data.Display.Article.Comments) &&
		len(data.Display.Article.Comments) > 6 {
		data.Display.IsMostLikes = true
	}

	return c.Render(http.StatusOK, "article", data)
}

func (h *Handler) ArticleLike(c echo.Context) (err error) {
	id := c.Param("id")
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path=article/"+id)
	}
	pos := c.QueryParam("pos")
	//Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	a := &model.Article{}
	if err = db.DB(MONGO_DB).C(ARTICLE).
		FindId(bson.ObjectIdHex(id)).One(a); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}
	if a.UserLiked == nil {
		a.UserLiked = make(map[string]bool)
	}
	if v, ok := a.UserLiked[data.ID]; ok {
		if v {
			a.Like--
			a.UserLiked[data.ID] = false
		} else {
			a.Like++
			a.UserLiked[data.ID] = true
		}
	} else {
		a.Like++
		a.UserLiked[data.ID] = true
	}

	if err = db.DB(MONGO_DB).C(ARTICLE).
		UpdateId(a.ID, a); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	return c.Redirect(http.StatusFound, "/article/"+id+"#"+pos)
}

func (h *Handler) RemoveArticle(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		log.Println("Not Login")
		return c.NoContent(http.StatusNotFound)
	}

	id := c.FormValue("articleid")
	if id == "" {
		id = c.Param("id")
	}

	db := h.DB.Clone()
	defer db.Close()
	//获取article
	a := &model.Article{}
	if err = db.DB(MONGO_DB).C(ARTICLE).
		FindId(bson.ObjectIdHex(id)).
		One(a); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	//权限检测
	if a.Editor != data.ID {
		admin := &model.User{}
		if err = db.DB(MONGO_DB).C(USER).
			FindId(bson.ObjectIdHex(data.ID)).
			One(admin); err != nil {
			if err == mgo.ErrNotFound {
				return echo.ErrNotFound
			} else {
				return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
			}
		}
		if !admin.Admin {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "您没有删除该文章的权限"}
		}
	}

	//删除
	if err = db.DB(MONGO_DB).C(ARTICLE).
		RemoveId(bson.ObjectIdHex(id)); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	//获取user
	u := &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(a.Editor)).
		One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	for i, v := range u.Articles {
		if v == id {
			u.Articles = append(u.Articles[:i], u.Articles[i+1:]...)
			break
		}
	}

	//更新user
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(a.Editor), u); err != nil {
		log.Printf("RemoveComment.Uodate User ERROR: %s", err)
	}

	//删除comments
	wg := sync.WaitGroup{}
	for i, v := range a.Comments {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
			if err := h.removeComment(v, true); err != nil {
				log.Printf("RemoveComment.WaitGroupGoroutine[%d] ERROR: %s", i, v)
			}
		}(i, v)
	}
	wg.Wait()

	//记录日志
	log := &model.Log{
		ID:        bson.NewObjectId(),
		Time:      time.Now(),
		Object:    a.ID.Hex(),
		Type:      "文稿",
		User:      data.ID,
		Operation: "删除",
	}

	if err = db.DB(MONGO_DB).C(LOG).
		Insert(log); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	return c.Redirect(http.StatusFound, "/user/dashboard")
}

func (h *Handler) ArticleImage(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}

	//Read file
	//Source
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("FormFile:")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "请先选择图片"}
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println()
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}
	defer src.Close()
	//Destination
	filePath := "static/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println("os.Create:")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("io.Copy")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	id := c.Param("id")
	url, err := toQiniu(id, file.Filename, filePath, "", "-imageStyle")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("os.Remove:")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	return c.Redirect(http.StatusFound, "/article/image?url="+url)
}

func (h *Handler) ArticleImageGet(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Url string
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}
	url := c.QueryParam("url")
	if url != "" {
		data.Url = url
	}
	return c.Render(http.StatusOK, "articlephoto", data)
}
