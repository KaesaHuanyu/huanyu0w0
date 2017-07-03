package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"log"
	"net/http"
	"sync"
	"time"
)

func (h *Handler) CreateComment(c echo.Context) (err error) {
	articleID := c.QueryParam("article")
	//获得当前登录账户的ID
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path=article/"+articleID)
	}

	comment := &model.Comment{
		ID:      bson.NewObjectId(),
		Time:    time.Now(),
		Article: articleID,
		Editor:  data.ID,
	}

	if err = c.Bind(comment); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}
	if comment.Content == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid content field"}
	}

	db := h.DB.Clone()
	defer db.Close()

	replyto := c.QueryParam("replyto")
	if replyto != "" {
		comment.Replyto = replyto
		//更新comment
		if err = db.DB(MONGO_DB).C(COMMENT).
			UpdateId(bson.ObjectIdHex(replyto), bson.M{"$addToSet": bson.M{"replies": comment.ID.Hex()}}); err != nil {
			if err == mgo.ErrNotFound {
				return echo.ErrNotFound
			} else {
				return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
			}
		}
	}

	if err = db.DB(MONGO_DB).C(COMMENT).Insert(comment); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	//更新user
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(comment.Editor), bson.M{"$addToSet": bson.M{"comments": comment.ID.Hex()}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}
	//更新article
	if err = db.DB(MONGO_DB).C(ARTICLE).
		UpdateId(bson.ObjectIdHex(comment.Article), bson.M{"$addToSet": bson.M{"comments": comment.ID.Hex()}}); err != nil {
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
		Object:    comment.ID.Hex(),
		Type:      "评论",
		User:      data.ID,
		Operation: "创建",
	}

	if err = db.DB(MONGO_DB).C(LOG).
		Insert(log); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	if replies := c.QueryParam("replies"); replies == "yes" {
		if comment := c.QueryParam("comment"); comment != "" {
			return c.Redirect(http.StatusFound, "/replies/"+comment)
		}
	}

	return c.Redirect(http.StatusFound, "/article/"+comment.Article)
}

func (h *Handler) CommentLike(c echo.Context) (err error) {
	article := c.QueryParam("article")
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path=article/"+article)
	}
	id := c.Param("id")
	pos := c.QueryParam("pos")
	//Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	comment := &model.Comment{}
	if err = db.DB(MONGO_DB).C(COMMENT).
		FindId(bson.ObjectIdHex(id)).One(comment); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}
	if comment.UserLiked == nil {
		comment.UserLiked = make(map[string]bool)
	}
	if v, ok := comment.UserLiked[data.ID]; ok {
		if v {
			comment.Like--
			comment.UserLiked[data.ID] = false
		} else {
			comment.Like++
			comment.UserLiked[data.ID] = true
		}
	} else {
		comment.Like++
		comment.UserLiked[data.ID] = true
	}

	if err = db.DB(MONGO_DB).C(COMMENT).
		UpdateId(comment.ID, comment); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	if replies := c.QueryParam("replies"); replies == "yes" {
		if comment := c.QueryParam("comment"); comment != "" {
			return c.Redirect(http.StatusFound, "/replies/"+comment+"#"+pos)
		}
	}

	return c.Redirect(http.StatusFound, "/article/"+article+"#"+pos)
}

func (h *Handler) Replies(c echo.Context) (err error) {
	//获得当前登录账户的ID
	data := &struct {
		model.Cookie
		DisplayComment *model.DisplayComment
		Replies        []*model.DisplayComment
		IsLike         bool
		IsFollow       bool
		IsEditor       bool
	}{
		DisplayComment: &model.DisplayComment{},
		Replies:        []*model.DisplayComment{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}

	id := c.Param("id")
	db := h.DB.Clone()
	defer db.Close()

	data.DisplayComment.Comment = &model.Comment{}
	if err = db.DB(MONGO_DB).C(COMMENT).
		FindId(bson.ObjectIdHex(id)).
		One(data.DisplayComment.Comment); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	data.DisplayComment.ID = data.DisplayComment.Comment.ID.Hex()
	data.DisplayComment.ShowTime = data.DisplayComment.Comment.GetShowTime()
	data.DisplayComment.ReplyNum = len(data.DisplayComment.Comment.Replies)

	data.DisplayComment.Editor = &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(data.DisplayComment.Comment.Editor)).
		One(data.DisplayComment.Editor); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}
	data.DisplayComment.Editor.Password = ""

	//是否为回复
	if data.DisplayComment.Comment.Replyto != "" {
		master := &model.Comment{}
		if err := db.DB(MONGO_DB).C(COMMENT).
			FindId(bson.ObjectIdHex(data.DisplayComment.Comment.Replyto)).
			One(master); err != nil {
			log.Println("<(￣︶￣)↗[GO!]", " Replyto:", err)
		}
		data.DisplayComment.Replyto = &model.User{}
		if err := db.DB(MONGO_DB).C(USER).
			FindId(bson.ObjectIdHex(master.Editor)).
			One(data.DisplayComment.Replyto); err != nil {
			log.Println("<(￣︶￣)↗[GO!]", " Replyto:", err)
		}
		data.DisplayComment.Replyto.Password = ""
	}

	//当前用户是否关注
	if v, ok := data.DisplayComment.Editor.IsFollower[data.ID]; ok {
		if v {
			data.DisplayComment.IsFollow = true
		}
	}

	//当前用户是否点赞
	if v, ok := data.DisplayComment.Comment.UserLiked[data.ID]; ok {
		if v {
			data.DisplayComment.IsLike = true
		}
	}

	wg := &sync.WaitGroup{}
	for i, v := range data.DisplayComment.Comment.Replies {
		wg.Add(1)
		data.Replies = append(data.Replies, &model.DisplayComment{})
		go func(i int, v string) {
			defer wg.Done()
			db := h.DB.Clone()
			defer db.Close()

			data.Replies[i].Number = i + 1
			data.Replies[i].Comment = &model.Comment{}
			if err = db.DB(MONGO_DB).C(COMMENT).
				FindId(bson.ObjectIdHex(v)).
				One(data.Replies[i].Comment); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, " Find comment:", err)
			}
			data.Replies[i].ID = data.Replies[i].Comment.ID.Hex()
			data.Replies[i].ShowTime = data.Replies[i].Comment.GetShowTime()
			data.Replies[i].ReplyNum = len(data.Replies[i].Comment.Replies)

			data.Replies[i].Editor = &model.User{}
			if err = db.DB(MONGO_DB).C(USER).
				FindId(bson.ObjectIdHex(data.Replies[i].Comment.Editor)).
				One(data.Replies[i].Editor); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, " Find editor:", err)
			}

			//是否是楼主
			if data.Replies[i].Comment.Editor == data.DisplayComment.Comment.Editor {
				data.Replies[i].IsEditor = true
			}

			//当前用户是否点赞
			if v, ok := data.Replies[i].Comment.UserLiked[data.ID]; ok {
				if v {
					data.Replies[i].IsLike = true
				}
			}

			//当前用户是否关注
			if v, ok := data.Replies[i].Editor.IsFollower[data.ID]; ok {
				if v {
					data.Replies[i].IsFollow = true
				}
			}

			data.Replies[i].Editor.Password = ""
		}(i, v)
	}
	wg.Wait()

	return c.Render(http.StatusOK, "replies", data)
}

func (h *Handler) RemoveComment(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		log.Println("Not Login")
		return c.NoContent(http.StatusNotFound)
	}

	id := c.FormValue("commentid")
	if id == "" {
		id = c.Param("id")
	}

	db := h.DB.Clone()
	defer db.Close()

	//获取comment
	comment := &model.Comment{}
	if err = db.DB(MONGO_DB).C(COMMENT).
		FindId(bson.ObjectIdHex(id)).
		One(comment); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	if comment.Editor != data.ID {
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
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "您没有删除该评论的权限"}
		}
	}

	if err = h.removeComment(id, false); err != nil {
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
		Object:    comment.ID.Hex(),
		Type:      "评论",
		User:      data.ID,
		Operation: "删除",
	}

	if err = db.DB(MONGO_DB).C(LOG).
		Insert(log); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	return c.Redirect(http.StatusFound, "/user/dashboard")
}

//删除评论
func (h *Handler) removeComment(id string, article bool) (err error) {
	db := h.DB.Clone()
	defer db.Close()

	//获取comment
	c := &model.Comment{}
	if err = db.DB(MONGO_DB).C(COMMENT).
		FindId(bson.ObjectIdHex(id)).
		One(c); err != nil {
		return
	}

	if c.Replyto != "" {
		r := &model.Comment{}
		if err = db.DB(MONGO_DB).C(COMMENT).
			FindId(bson.ObjectIdHex(c.Replyto)).
			One(r); err != nil {
			return
		}

		for i, v := range r.Replies {
			if v == id {
				r.Replies = append(r.Replies[:i], r.Replies[i+1:]...)
				break
			}
		}

		//更新user
		if err = db.DB(MONGO_DB).C(COMMENT).
			UpdateId(bson.ObjectIdHex(c.Replyto), r); err != nil {
			return
		}
	}

	for _, v := range c.Replies {
		if err = db.DB(MONGO_DB).C(COMMENT).
			UpdateId(bson.ObjectIdHex(v), bson.M{"$set": bson.M{"replyto": ""}}); err != nil {
			return
		}
	}

	//获取user
	u := &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(c.Editor)).
		One(u); err != nil {
		return
	}

	for i, v := range u.Comments {
		if v == id {
			u.Comments = append(u.Comments[:i], u.Comments[i+1:]...)
			break
		}
	}

	//更新user
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(c.Editor), u); err != nil {
		return
	}

	if !article {
		//获取article
		a := &model.Article{}
		if err = db.DB(MONGO_DB).C(ARTICLE).
			FindId(bson.ObjectIdHex(c.Article)).
			One(a); err != nil {
			return
		}

		for i, v := range a.Comments {
			if v == id {
				a.Comments = append(a.Comments[:i], a.Comments[i+1:]...)
				break
			}
		}

		//更新article
		if err = db.DB(MONGO_DB).C(ARTICLE).
			UpdateId(bson.ObjectIdHex(c.Article), a); err != nil {
			return
		}
	}

	//删除comment
	if err = db.DB(MONGO_DB).C(COMMENT).
		RemoveId(bson.ObjectIdHex(id)); err != nil {
		return
	}

	return
}
