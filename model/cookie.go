package model

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type (
	Cookie struct {
		ID       string `json:"id,omitempty" bson:"id,omitempty"`
		Avatar   string `json:"avatar,omitempty" bson:"avatar,omitempty"`
		Remember bool
		IsLogin  bool
	}
)

func (cookie *Cookie) WriteCookie(c echo.Context) {
	t := 24 * time.Hour
	if cookie.Remember {
		t = 30 * 24 * time.Hour
	}
	cookie1 := new(http.Cookie)
	cookie1.Expires = time.Now().Add(t)
	cookie1.Name = "id"
	cookie1.Value = cookie.ID
	c.SetCookie(cookie1)

	cookie2 := new(http.Cookie)
	cookie2.Expires = time.Now().Add(t)
	cookie2.Name = "avatar"
	cookie2.Value = cookie.Avatar
	c.SetCookie(cookie2)
}

func (cookie *Cookie) ReadCookie(c echo.Context) (err error) {
	ck, err := c.Cookie("id")
	if err != nil {
		return
	}
	cookie.ID = ck.Value

	ck, err = c.Cookie("avatar")
	if err != nil {
		return
	}
	cookie.Avatar = ck.Value
	return nil
}
