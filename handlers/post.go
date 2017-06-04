package handlers

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}
	u := &model.User{
		ID: bson.ObjectIdHex(data.ID),
	}
	p := &model.Post{
		ID:   bson.NewObjectId(),
		From: u.ID.Hex(),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	//Validate
	if p.To == "" || p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid to or message field"}
	}

	//Find user from database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("huanyu0w0").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	//Save post in database
	if err = db.DB("huanyu0w0").C("posts").Insert(p); err != nil {
		return
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handler) FetchPost(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	//Default
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	posts := []*model.Post{}
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("huanyu0w0").C("posts").
		Find(bson.M{"to": data.ID}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&posts); err != nil {
		return
	}

	return c.JSON(http.StatusOK, posts)
}
