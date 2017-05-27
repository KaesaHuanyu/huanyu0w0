package controllers

import (
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"log"
	"net/http"
	"time"
)

func SignupGet(c echo.Context) error {
	data := new(model.Cookies)
	isSignin, err := c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ", err)
	} else if isSignin.Value == "yes" {
		data.IsLogin = true
		userid, _ := c.Cookie("userid")
		data.UserId = userid.Value
		avatar, _ := c.Cookie("avatar")
		data.Avatar = avatar.Value
	}
	return c.Render(http.StatusOK, "signup", data)
}

func SigninGet(c echo.Context) error {
	data := new(model.Cookies)
	isSignin, err := c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ", err)
	} else if isSignin.Value == "yes" {
		data.IsLogin = true
		userid, err := c.Cookie("userid")
		if err == nil {
			data.UserId = userid.Value
		}
		return c.Redirect(http.StatusFound, "/user/"+userid.Value)
	}
	return c.Render(http.StatusOK, "signin", data)
}

func SigninPost(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")
	remember := c.FormValue("remember")
	//登录验证
	if len(email) == 0 || len(password) == 0 {
		return c.Render(http.StatusFound, "error", "请输入用户名和密码")
	}
	u := new(model.User)
	err := model.FindOneNoId(model.MONGO_USER, "email", email, u)
	if err != nil {
		log.Println("Not found user: ", email)
		return c.Render(http.StatusFound, "error", "用户名不存在")
	}
	if u.Password != password {
		log.Println("Password wrong.")
		return c.Render(http.StatusFound, "error", "密码错误")
	}
	//Set Cookies
	isSignin := new(http.Cookie)
	isSignin.Name = "isSignin"
	isSignin.Value = "yes"
	if remember == "on" {
		isSignin.Expires = time.Now().Add(model.REMEMBER)
	} else {
		isSignin.Expires = time.Now().Add(model.COOKIE_TIME)
	}
	c.SetCookie(isSignin)

	userid := new(http.Cookie)
	userid.Name = "userid"
	userid.Value = u.ID.Hex()
	if remember == "on" {
		userid.Expires = time.Now().Add(model.REMEMBER)
	} else {
		userid.Expires = time.Now().Add(model.COOKIE_TIME)
	}
	c.SetCookie(userid)

	avatar := new(http.Cookie)
	avatar.Name = "avatar"
	avatar.Value = u.Avatar
	if remember == "on" {
		userid.Expires = time.Now().Add(model.REMEMBER)
	} else {
		userid.Expires = time.Now().Add(model.COOKIE_TIME)
	}
	c.SetCookie(avatar)

	like := new(http.Cookie)
	like.Name = "like"
	like.Value = "false"
	if remember == "on" {
		like.Expires = time.Now().Add(model.REMEMBER)
	} else {
		like.Expires = time.Now().Add(model.COOKIE_TIME)
	}
	c.SetCookie(like)

	log.Println("sign in successed.")
	return c.Redirect(http.StatusFound, "/")
}

func Signout(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}
	return c.Redirect(http.StatusFound, "/signin")
}
