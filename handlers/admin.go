package handlers

import (
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func (h *Handler) Admin(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Admin *model.User
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

	if !data.Admin.Admin {
		log.Println("Not Admin")
		return c.NoContent(http.StatusNotFound)
	}
	return c.Render(http.StatusOK, "admin", data)
}
