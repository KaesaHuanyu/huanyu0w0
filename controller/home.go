package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"huanyu0w0/model"
	"log"
)

func Home(c echo.Context) error {
	data := &struct {
		model.Cookies
		articles []*model.Article
	}{}
	isSignin, err := c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ", err)
		return c.Redirect(http.StatusFound, "/signin")
	} else if isSignin.Value == "yes" {
		data.IsLogin = true
		userid, err := c.Cookie("userid")
		if err != nil {
			return c.Render(http.StatusFound, "error", err)
		}
		data.UserId = userid.Value
		avatar, err := c.Cookie("avatar")
		if err != nil {
			return c.Render(http.StatusFound, "error", err)
		}
		data.Avatar = avatar.Value
	}

	err = model.FindAll(model.MONGO_ARTICLE, "", "", 20, "like", data.articles)
	if err != nil {
		log.Println("Main FindAll error.")
		return c.Render(http.StatusFound, "error", "出了点小问题...")
	}
	return c.Render(http.StatusOK, "home", data)
}
