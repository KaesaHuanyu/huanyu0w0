package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"huanyu0w0/model"
	"log"
	"html/template"
	"github.com/russross/blackfriday"
)

func Home(c echo.Context) error {
	data := &struct {
		model.Cookies
		ArticleEditors []*model.ArticleEditor
	}{
		ArticleEditors: []*model.ArticleEditor{},
	}

	isSignin, err := c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ", err)
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
	articles := &[]*model.Article{}
	err = model.FindAll(model.MONGO_ARTICLE, "", "", 5, "-like", articles)
	if err != nil {
		log.Println("Home FindAll error.", err)
		return c.Render(http.StatusFound, "error", "出了点小问题...")
	}
	for index, article := range *articles {
		data.ArticleEditors = append(data.ArticleEditors, &model.ArticleEditor{
			&model.Article{},
			&model.User{},
			"",
		})
		data.ArticleEditors[index].Article = article
		data.ArticleEditors[index].Introduction = template.HTML(blackfriday.MarkdownCommon([]byte(article.Introduction)))
		err = model.FindMongo(model.MONGO_USER, "_id", article.Editor, data.ArticleEditors[index].Editor)
		if err != nil {
			log.Println("Home FindMongo editor error.", err)
			//return c.Render(http.StatusFound, "error", "出了点小问题...")
		}
	}
	log.Println(data)
	return c.Render(http.StatusOK, "home", data)
}
