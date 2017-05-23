//docker run --name huayu0w0-mongo -d --restart always -v ~/mongo:/data/db mongo
//docker run -d --name huanyu0w0-server-test --restart always --link huanyu0w0-mongo-test:mongo -p 80:1323 daocloud.io/kaesa/huanyu0w0-server-test
package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"huanyu0w0/controller"
)

//实现echo.Renderer接口
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	//声明模版集
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	//Echo instance
	e := echo.New()

	//注册模版
	e.Renderer = t

	//Echo router
	e.GET("/", controller.Home)

	e.GET("/signup", controller.SignupGet)
	e.POST("/signup", controller.CreateUser)

	e.GET("/signin", controller.SigninGet)
	e.POST("/signin", controller.SigninPost)

	e.GET("/signout", controller.Signout)

	//CRUD Article
	e.GET("/article", controller.Article)
	e.GET("/article/create", controller.CreateArticleGet)
	e.POST("/article/create", controller.CreateArticlePost)
	e.GET("/article/:id", controller.GetArticle)
	e.PUT("/article/:id", controller.UpdateArticle)
	e.DELETE("/article/:id", controller.DeleteArticle)
	e.GET("/like/:id", controller.LikeArticle)

	//CRUD User
	//e.POST("/user", controller.CreateUser)
	e.GET("/user/:id", controller.GetUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUser)

	//CRD Comment
	e.POST("/comment", controller.CreateComment)
	e.GET("/comment/:id", controller.GetComment)
	e.DELETE("/comment/:id", controller.DeleteComment)
	e.GET("/nice/:id", controller.NiceComment)
	e.GET("/reply/:id", controller.ReplyComment)

	e.GET("/curriculumVitae", controller.CurriculumVitae)

	//图片之类的静态文件路由
	e.Static("/static", "static")
	//e.File("/", "view/article.html")
	//运行服务器
	e.Logger.Printf("%s\n", e.Start(":1323"))
}
