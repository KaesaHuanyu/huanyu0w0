package model

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type (
	Cookie struct {
		ID       string
		Avatar   string
		Email	string
		UnixTime string
		Name	string
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

	cookie3 := new(http.Cookie)
	cookie3.Expires = time.Now().Add(t)
	cookie3.Name = "email"
	cookie3.Value = cookie.Email
	c.SetCookie(cookie3)

	cookie4 := new(http.Cookie)
	cookie4.Expires = time.Now().Add(t)
	cookie4.Name = "time"
	cookie4.Value = cookie.UnixTime
	c.SetCookie(cookie4)

	cookie5 := new(http.Cookie)
	cookie5.Expires = time.Now().Add(t)
	cookie5.Name = "name"
	cookie5.Value = cookie.Name
	c.SetCookie(cookie5)
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

	ck, err = c.Cookie("email")
	if err != nil {
		return
	}
	cookie.Email = ck.Value

	ck, err = c.Cookie("time")
	if err != nil {
		return
	}
	cookie.UnixTime = ck.Value

	ck, err = c.Cookie("name")
	if err != nil {
		return
	}
	cookie.Name = ck.Value
	return nil
}
