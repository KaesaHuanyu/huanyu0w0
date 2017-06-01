//docker run --name huayu0w0-mongo -d --restart always -v ~/mongo:/data/db mongo
//docker run -d --name huanyu0w0-server-test --restart always --link huanyu0w0-mongo-test:mongo -p 80:1323 daocloud.io/kaesa/huanyu0w0-server-test
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"html/template"
	"huanyu0w0/handlers"
	"io"
)

//实现echo.Renderer接口
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	//Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	//e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey: []byte(handlers.Key),
	//	Skipper: func(c echo.Context) bool {
	//		//Skip authentication for those requests
	//		if c.Path() == "/login" || c.Path() == "/signup" || c.Path() == "/" || c.Path() == "/static"{
	//			return true
	//		}
	//		return false
	//	},
	//}))

	//声明模版集
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	//注册模版
	e.Renderer = t

	//Database connection
	db, err := mgo.Dial(handlers.MONGO_ADDRESS)
	if err != nil {
		e.Logger.Fatal(err)
	}

	//Create indices
	if err = db.Copy().DB(handlers.MONGO_DB).C(handlers.USER).EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	if err = db.Copy().DB(handlers.MONGO_DB).C(handlers.ARTICLE).EnsureIndex(mgo.Index{
		Key:    []string{"topic", "like"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	//Initialize
	h := &handlers.Handler{db}
	//Routes
	e.GET("/", h.Home)
	e.GET("/signup", h.SignupGet)
	e.POST("/signup", h.Signup)
	e.GET("/login", h.Signin)
	e.POST("/login", h.Login)
	e.GET("/signout", h.Signout)
	e.GET("/article/create", h.CreateArticleGet)
	e.POST("/article/create", h.CreateArticle)
	e.GET("/article/:id", h.ArticleDetail)
	e.GET("/user/:id", h.UserDetail)
	e.GET("/topic/:topic", h.Topic)
	e.POST("/follow/:id", h.Follow)
	e.POST("/posts", h.CreatePost)
	e.GET("/feed", h.FetchPost)
	e.GET("/curriculumVitae", h.CurriculumVitae)
	//图片之类的静态文件路由
	e.Static("/static", "static")

	//Run
	e.Logger.Fatal(e.Start(":1323"))
}
