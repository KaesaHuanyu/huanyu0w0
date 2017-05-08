package controller

import (
	"github.com/russross/blackfriday"
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"net/http"
	"time"
	"html/template"
	"github.com/satori/go.uuid"
	"log"
)

func Article(c echo.Context) error {
	data := struct {
		isSignin bool
	}{
		isSignin: false,
	}
	_, err := c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ", err)
	} else {
		data.isSignin = true
	}
	return c.Render(http.StatusOK, "article", data)
}

func CreateArticleGet(c echo.Context) error {
	data := new(model.Cookies)
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

	return c.Render(http.StatusOK, "createarticle", data)
}

func CreateArticlePost(c echo.Context) error {
	u := new(model.User)
	userid, err := c.Cookie("userid")
	if err != nil {
		return c.Render(http.StatusFound, "error", err)
	}
	err = model.FindMongo(model.MONGO_USER, "_id", userid.Value, u)
	if err != nil {
		return c.Render(http.StatusFound, "error", err)
	}

	a := &model.Article{
		Id: uuid.NewV4().String(),
		Time: time.Now(),
		Editor: u.Id,
	}
	if err := c.Bind(a); err != nil {
		log.Println("CreateArticle context.Bind() error.")
		return c.Render(http.StatusFound, "error", err)
	}
	if len(a.Content) >= 20 {
		a.Introduction = a.Content[:20] + "..."
	} else {
		a.Introduction = a.Content
	}
	//User中插入article
	u.Articles = append(u.Articles, a.Id)
	err = model.UpdateMongo(model.MONGO_USER, u.Id, u)
	if err != nil {
		log.Println("CreateArticle model.UpdateMongo() error.", err)
		return c.Render(http.StatusFound, "error", "保存安利💊失败...")
	}
	//插入article到数据库
	err = model.InsertMongo(model.MONGO_ARTICLE, a)
	if err != nil {
		log.Println("CreateArticle model.InsertMongo() error.", err)
		return c.Render(http.StatusFound, "error", "保存安利💊失败...")
	}
	log.Println(a)
	return c.Redirect(http.StatusFound, "/article/" + a.Id)
}

func GetArticle(c echo.Context) error {
	data := &struct {
		model.Cookies
		Article *model.Article
		Editor *model.User
		Comments []*model.CommentEditor
		Content template.HTML
	}{
		Article: &model.Article{},
		Editor: &model.User{},
		Comments: []*model.CommentEditor{},
	}

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

	id := c.Param("id")
	err = model.FindMongo(model.MONGO_ARTICLE, "_id" , id, data.Article)
	if err != nil {
		log.Println("GetArticle FindMongo Article error: ",err)
		return c.Render(http.StatusFound, "error", "404 NOT FOUND")
	}
	data.Content = template.HTML(blackfriday.MarkdownCommon([]byte(data.Article.Content)))
	err = model.FindMongo(model.MONGO_USER, "_id", data.Article.Editor, data.Editor)
	if err != nil {
		log.Println("GetArticle FindMongo User error: ", err)
		return c.Render(http.StatusFound, "error", "404 NOT FOUND")
	}
	for index, comment := range data.Article.Comments {
		model.FindMongo(model.MONGO_COMMENT, "_id", comment, data.Comments[index].Comment)
		model.FindMongo(model.MONGO_USER, "_id", data.Comments[index].Comment.Editor, data.Comments[index].Editor)
	}
	return c.Render(http.StatusOK, "article", data)
}

func UpdateArticle(c echo.Context) error {
	a := &model.Article{
		Change: time.Now(),
	}
	if err := c.Bind(a); err != nil {
		log.Println("UpdateArticle Bind error.")
		return err
	}
	id := c.Param("id")
	err := model.UpdateMongo(model.MONGO_ARTICLE, id, a)
	if err != nil {
		log.Println("UpdateArticle UpdateMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func DeleteArticle(c echo.Context) error {
	id := c.Param("id")
	err := model.RemoveMongo(model.MONGO_ARTICLE, id)
	if err != nil {
		log.Println("DeleteArticle RemoveMongo error.")
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func Like(c echo.Context) error {

	_, err := c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ERROR: ", err)
		return c.Redirect(http.StatusFound, "/signin")
	}

	userid, err := c.Cookie("userid")
	if err != nil {
		return c.Redirect(http.StatusFound, "/signin")
	}

	id := c.Param("id")
	a := new(model.Article)
	err = model.FindMongo(model.MONGO_ARTICLE, "_id", id, a)
	if err != nil {
		log.Println("GetArticle FindMongo error.")
		return c.Render(http.StatusNotFound, "error", "404 not found")
	}
	if a.UserLiked[userid.Value] == true {
		return c.Render(http.StatusFound, "error", "您已经被安立💊过了~")
	}
	a.Like++
	a.UserLiked[userid.Value] = true
	err = model.UpdateMongo(model.MONGO_ARTICLE, id, a)
	if err != nil {
		log.Println("Update error: ", err)
		c.Render(http.StatusFound, "error", "牙白...出了点问题...")
	}
	return c.Redirect(http.StatusFound, "/article/" + id)
}