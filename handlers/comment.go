package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"huanyu0w0/model"
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

func (h *Handler) CreateComment(c echo.Context) (err error) {
	//获得当前登录账户的ID
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}

	articleID := c.QueryParam("article")
	comment := &model.Comment{
		ID: bson.NewObjectId(),
		Time: time.Now(),
		Article: articleID,
		Editor: data.ID,
	}
	if err = c.Bind(comment); err != nil {
		return
	}
	if comment.Content == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid content field"}
	}

	db := h.DB.Clone()
	defer db.Close()
	//更新user
	if err = db.DB(MONGO_DB).C(USER).
	UpdateId(bson.ObjectIdHex(comment.Editor), bson.M{"$addToSet": bson.M{"comments": comment.ID.Hex()}}); err != nil{
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return err
		}
	}
	//更新article
	if err = db.DB(MONGO_DB).C(ARTICLE).
		UpdateId(bson.ObjectIdHex(comment.Article), bson.M{"$addToSet": bson.M{"comments": comment.ID.Hex()}}); err != nil{
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return err
		}
	}

	if err = db.DB(MONGO_DB).C(COMMENT).Insert(comment); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}
