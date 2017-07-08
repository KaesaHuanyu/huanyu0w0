//docker run --name huayu0w0-mongo -d --restart always -m 1024m --oom-kill-disable -v ~/mongo:/data/db mongo
//docker run -d --name huanyu0w0-server --restart always --link huanyu0w0-mongo:mongo -p 80:1323 daocloud.io/kaesa/huanyu0w0-server:v1.0.1
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"html/template"
	"huanyu0w0/handlers"
	"io"
	"github.com/afocus/captcha"
	"image/png"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"net/http"
)

//实现echo.Renderer接口
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var cap *captcha.Captcha

func main() {
	//Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS中间件
	e.Use(middleware.CORS())
	//session中间件
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))


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

	if err = db.Copy().DB(handlers.MONGO_DB).C(handlers.USER).EnsureIndex(mgo.Index{
		Key:    []string{"name"},
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
	e.GET("/articlelike/:id", h.ArticleLike)
	e.GET("/commentlike/:id", h.CommentLike)
	e.POST("/comment/create", h.CreateComment)
	e.GET("/replies/:id", h.Replies)
	e.GET("/user/:id", h.UserDetail)
	e.GET("/user/dashboard", h.Dashboard)
	e.GET("/admin/dashboard", h.Admin)
	e.GET("/topic/:topic", h.Topic)
	e.GET("/follow/:id", h.Follow)
	e.POST("/posts", h.CreatePost)
	e.GET("/feed", h.FetchPost)
	e.POST("/updateAvatar", h.UpdateAvatar)
	e.GET("/resume", h.CurriculumVitae)
	e.GET("/thankYouForYourGenerosity", h.ThankYou)
	e.GET("/article/remove/:id", h.RemoveArticle)
	e.GET("/comment/remove/:id", h.RemoveComment)
	//e.GET("/user/remove/:id", h.RemoveUser)
	e.POST("/article/remove", h.RemoveArticle)
	e.POST("/comment/remove", h.RemoveComment)
	e.POST("/user/update", h.UpdateUser)
	e.GET("/logs", h.AdminLogs)
	e.GET("/article/image", h.ArticleImageGet)
	e.POST("/article/image/:id", h.ArticleImage)
	//图片之类的静态文件路由
	e.Static("/static", "static")
	//验证码
	cap = captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err)
	}
	cap.SetSize(160, 80)
	e.GET("/verify", func(c echo.Context) (err error) {
		img, str := cap.Create(4, captcha.NUM)
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   120,
			HttpOnly: true,
		}
		sess.Values["verify"] = str
		sess.Save(c.Request(), c.Response())
		defer png.Encode(c.Response().Writer, img)
		return c.NoContent(http.StatusOK)
	})

	//Run
	e.Logger.Fatal(e.Start(":1323"))
}
