package handlers

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"net/http"
	"time"
	"sync"
	"fmt"
)

func (h *Handler) CreateArticleGet(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
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
		return
	}

	if a.Reason == "" || a.Title == "" || a.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid title, name or reason field"}
	}

	db := h.DB.Clone()
	defer db.Close()
	//更新user的articles
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$addToSet": bson.M{"articles": a.ID.Hex()}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	if err = db.DB(MONGO_DB).C(ARTICLE).Insert(a); err != nil {
		return
	}

	return c.JSON(http.StatusOK, a)
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

	data.Display.ID = data.Display.Article.ID.Hex()
	data.Display.ShowTime = data.Display.Article.GetShowTime()
	data.Display.ShowTopic = data.Display.Article.GetShowTopic()

	data.Display.Editor = &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
	FindId(bson.ObjectIdHex(data.Display.Article.Editor)).
	One(data.Display.Editor); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	wg := sync.WaitGroup{}
	for i, v := range data.Display.Article.Comments {
		wg.Add(1)
		data.Display.Comments = append(data.Display.Comments, &model.DisplayComment{})
		go func(i int, v string) {
			defer wg.Done()

			data.Display.Comments[i].ID = data.Display.Comments[i].Comment.ID.Hex()
			data.Display.Comments[i].ShowTime = data.Display.Comments[i].Comment.GetShowTime()
			data.Display.Comments[i].ReplyNum = len(data.Display.Comments[i].Comment.Replies)
			db := h.DB.Clone()
			defer db.Close()

			data.Display.Comments[i].Comment = &model.Comment{}
			if err := db.DB(MONGO_DB).C(COMMENT).
			FindId(bson.ObjectIdHex(v)).
			One(data.Display.Comments[i].Comment); err != nil {
				fmt.Println("<(￣︶￣)↗[GO!]", i, ":", err)
			}

			data.Display.Comments[i].Editor = &model.User{}
			if err := db.DB(MONGO_DB).C(USER).
				FindId(bson.ObjectIdHex(data.Display.Comments[i].Comment.Editor)).
				One(data.Display.Comments[i].Editor); err != nil {
				fmt.Println("<(￣︶￣)↗[GO!]", i, ":", err)
			}
		}(i, v)
	}
	wg.Wait()

	return c.Render(http.StatusOK, "article", data)
}