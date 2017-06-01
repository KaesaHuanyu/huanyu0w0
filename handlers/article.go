package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"huanyu0w0/model"
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
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
		ID: bson.NewObjectId(),
		Time: time.Now(),
		Editor: data.ID,
	}

	if err = c.Bind(a); err != nil {
		return
	}

	if a.Reason == "" || a.Title == "" || a.Name == ""{
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