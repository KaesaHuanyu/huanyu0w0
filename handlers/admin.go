package handlers

import (
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

func (h *Handler) Admin(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Admin *model.User
		User int
		Article int
		Comment int
	}{
		Admin: &model.User{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		log.Println("Not Login")
		return c.NoContent(http.StatusNotFound)
	}

	//取得mongo连接
	db := h.DB.Clone()
	defer db.Close()

	//得到User
	data.Admin = &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(data.ID)).
		One(data.Admin); err != nil{
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return c.NoContent(http.StatusNotFound)
		}
	}

	data.Admin.Password = ""

	if !data.Admin.Admin {
		log.Println("Not Admin")
		return c.NoContent(http.StatusNotFound)
	}

	data.User, _ = db.DB(MONGO_DB).C(USER).Find(nil).Count()
	data.Article, _ = db.DB(MONGO_DB).C(ARTICLE).Find(nil).Count()
	data.Comment, _ = db.DB(MONGO_DB).C(COMMENT).Find(nil).Count()

	return c.Render(http.StatusOK, "admin", data)
}

func (h *Handler) AdminLogs(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Admin *model.User
		Events []*struct{
			Time string
			User string
			Admin bool
			Operation string
			Type string
			Object string
			Signup bool
		}
	}{
		Admin: &model.User{},
		Events: []*struct{
			Time string
			User string
			Admin bool
			Operation string
			Type string
			Object string
			Signup bool
		}{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		log.Println("Not Login")
		return c.NoContent(http.StatusNotFound)
	}

	//取得mongo连接
	db := h.DB.Clone()
	defer db.Close()

	//得到User
	data.Admin = &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(data.ID)).
		One(data.Admin); err != nil{
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return c.NoContent(http.StatusNotFound)
		}
	}

	data.Admin.Password = ""

	if !data.Admin.Admin {
		log.Println("Not Admin")
		return c.NoContent(http.StatusNotFound)
	}

	//读取日志
	logs := []*model.Log{}
	if err = db.DB(MONGO_DB).C(LOG).
	Find(nil).Sort("-time").
	All(&logs); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	for _, v := range logs {
		u := &model.User{}
		if err := db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(v.User)).
		One(u); err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
		event := &struct{
			Time string
			User string
			Admin bool
			Operation string
			Type string
			Object string
			Signup bool
		}{
			Time: v.Time.Format("2006年 01月02日 15:04:05"),
			User: u.Name,
			Admin: u.Admin,
			Operation: v.Operation,
			Type: v.Type,
			Object: v.Object,
			Signup: v.Signup,
		}
		data.Events = append(data.Events, event)
	}

	return c.Render(http.StatusOK, "logs", data)
}
