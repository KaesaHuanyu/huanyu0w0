//docker run --name huayu0w0-mongo -d --restart always -v ~/mongo:/data/db mongo
//docker run -d --name huanyu0w0-server-test --restart always --link huanyu0w0-mongo-test:mongo -p 80:1323 daocloud.io/kaesa/huanyu0w0-server-test
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/cookbook/twitter/handler"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"html/template"
	"huanyu0w0/controllers"
	"huanyu0w0/model"
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
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		//Skipper: func(c echo.Context) bool {
		//	//Skip authentication for signup and login requests
		//	if c.Path() == "/Login" || c.Path() == "/Signup" || c.Path() == "/"{
		//		return true
		//	}
		//	return false
		//},
		Skipper: func(c echo.Context) bool {
			return true
		},
	}))

	//声明模版集
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	//注册模版
	e.Renderer = t

	//Database connection
	db, err := mgo.Dial(model.MONGO_ADDRESS)
	if err != nil {
		e.Logger.Fatal(err)
	}

	//Create indices
	if err = db.Copy().DB("huanyu0w0").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	//Initialize
	h := &handler.Handler{db}
	//Routes
	e.GET("/signup", h.SignupGet)
	e.POST("/signup", h.Signup)
	e.GET("/login", h.Signin)
	e.POST("/login", h.Login)
	e.POST("/follow/:id", h.Follow)
	e.POST("/posts", h.CreatePost)
	e.GET("/feed", h.FetchPost)

	//Echo router
	e.GET("/", controllers.Home)

	e.GET("/signup", controllers.SignupGet)
	//e.POST("/signup", controllers.CreateUser)

	e.GET("/Login", controllers.SigninGet)
	e.POST("/signin", controllers.SigninPost)

	e.GET("/signout", controllers.Signout)

	//CRUD Article
	e.GET("/article", controllers.Article)
	e.GET("/article/create", controllers.CreateArticleGet)
	e.POST("/article/create", controllers.CreateArticlePost)
	e.GET("/article/:id", controllers.GetArticle)
	e.PUT("/article/:id", controllers.UpdateArticle)
	e.DELETE("/article/:id", controllers.DeleteArticle)
	e.GET("/like/:id", controllers.LikeArticle)

	//CRUD User
	//e.POST("/user", controllers.CreateUser)
	e.GET("/user/:id", controllers.GetUser)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	//CRD Comment
	e.POST("/comment", controllers.CreateComment)
	e.GET("/comment/:id", controllers.GetComment)
	e.DELETE("/comment/:id", controllers.DeleteComment)
	e.GET("/nice/:id", controllers.NiceComment)
	e.GET("/reply/:id", controllers.ReplyComment)

	e.GET("/curriculumVitae", controllers.CurriculumVitae)

	//图片之类的静态文件路由
	e.Static("/static", "static")
	//e.File("/", "view/article.html")
	//运行服务器
	//e.Logger.Printf("%s\n", e.Start(":1323"))
	e.Logger.Fatal(e.Start(":1323"))
}
