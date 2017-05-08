package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"log"
	"huanyu0w0/model"
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
		//username, _ := c.Cookie("username")
		//data.UserName = username.Value
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
		avatar, err := c.Cookie("avatar")
		if err == nil {
			data.Avatar = avatar.Value
		}
		//username, err := c.Cookie("username")
		//if err == nil {
		//	data.UserName = username.Value
		//}
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
	err := model.FindMongo(model.MONGO_USER, "email", email, u)
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
	userid.Value = u.Id
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

	//username := new(http.Cookie)
	//username.Name = "username"
	//username.Value = u.Name
	//username.Expires = time.Now().Add(3 * time.Hour)
	//c.SetCookie(username)
	log.Println("sign in successed.")
	return c.Redirect(http.StatusFound, "/signup")
}

func Signout(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}
	return c.Redirect(http.StatusFound, "/signin")
}