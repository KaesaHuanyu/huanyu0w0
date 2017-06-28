package handlers

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"log"
	"net/http"
	"sync"
	"time"
)

func (h *Handler) CreateArticleGet(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
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
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	if a.Reason == "" || a.Title == "" || a.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid title, name or reason field"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB(MONGO_DB).C(ARTICLE).Insert(a); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	//更新user的articles
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$addToSet": bson.M{"articles": a.ID.Hex()}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
		}
	}

	return c.Redirect(http.StatusFound, "/article/"+a.ID.Hex())
}

func (h *Handler) ArticleDetail(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Display *model.Display
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

	data.Display.Editor.Password = ""
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
			return err
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
			return err
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
	return c.Redirect(http.StatusFound, "/user/dashboard")
}