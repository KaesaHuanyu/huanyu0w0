package controllers

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"log"
	"net/http"
	"time"
)

func CreateComment(c echo.Context) error {
	article := c.FormValue("article")
	content := c.FormValue("content")
	editor, err := c.Cookie("userid")
	if err != nil {
		log.Println("not sign in, ", err)
		return c.Redirect(http.StatusFound, "/signin")
	}

	comment := &model.Comment{
		ID:      bson.NewObjectId(),
		Time:    time.Now(),
		Article: bson.ObjectIdHex(article),
		Editor:  bson.ObjectIdHex(editor.Value),
		Content: content,
	}

	if err := c.Bind(comment); err != nil {
		log.Println("CreateComment Bind error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	err = model.Insert(model.MONGO_COMMENT, comment)
	if err != nil {
		log.Println("CreateComment InsertMongo error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}

	a := new(model.Article)
	err = model.FindOne(model.MONGO_ARTICLE, bson.ObjectIdHex(article), a)
	if err != nil {
		log.Println("CreateComment FindMongo article error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	a.Comments = append(a.Comments, comment.ID)
	err = model.Update(model.MONGO_ARTICLE, bson.ObjectIdHex(article), a)
	if err != nil {
		log.Println("CreateComment UpdateMongo article error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}

	u := new(model.User)
	err = model.FindOne(model.MONGO_USER, bson.ObjectIdHex(editor.Value), u)
	if err != nil {
		log.Println("CreateComment FindMongo editor error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	u.Comments = append(u.Comments, comment.ID)
	err = model.Update(model.MONGO_USER, bson.ObjectIdHex(editor.Value), u)
	if err != nil {
		log.Println("CreateComment UpdateMongo editor error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	return c.Redirect(http.StatusFound, "/article/"+comment.Article.Hex())
}

func GetComment(c echo.Context) error {
	id := c.Param("id")
	comment := new(model.Comment)
	err := model.FindOne(model.MONGO_COMMENT, bson.ObjectIdHex(id), comment)
	if err != nil {
		log.Println("GetComment FindMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

func DeleteComment(c echo.Context) error {
	id := c.Param("id")
	err := model.Remove(model.MONGO_COMMENT, bson.ObjectIdHex(id))
	if err != nil {
		log.Println("DeleteComment RemoveMongo error.")
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func NiceComment(c echo.Context) error {
	id := c.Param("id")
	user, err := c.Cookie("userid")
	if err != nil {
		log.Println("not sign in, ", err)
		return c.Redirect(http.StatusFound, "/signin")
	}
	nice := new(http.Cookie)
	nice.Name = id

	comment := new(model.Comment)
	err = model.FindOne(model.MONGO_COMMENT, bson.ObjectIdHex(id), comment)
	if err != nil {
		log.Println("NiceComment FindMongo error, ", err)
		return c.Render(http.StatusFound, "error", "赞失败了...找不到这条评论...")
	}
	if comment.UserLiked[bson.ObjectIdHex(user.Value)] == true {
		comment.Like--
		comment.UserLiked[bson.ObjectIdHex(user.Value)] = false
		nice.Value = "false"
		c.SetCookie(nice)
		err := model.Update(model.MONGO_COMMENT, bson.ObjectIdHex(id), comment)
		if err != nil {
			log.Println("NiceComment Update error: ", err)
			c.Render(http.StatusFound, "error", "牙白...出了点问题...")
		}
		return c.Redirect(http.StatusFound, "/article/"+comment.Article.Hex()+"#"+comment.ID.Hex())
	}
	comment.Like++
	comment.UserLiked[bson.ObjectIdHex(user.Value)] = true
	nice.Value = "true"
	err = model.Update(model.MONGO_COMMENT, bson.ObjectIdHex(id), comment)
	if err != nil {
		log.Println("NiceComment Update error: ", err)
		c.Render(http.StatusFound, "error", "牙白...出了点问题...")
	}
	return c.Redirect(http.StatusFound, "/article/"+comment.Article.Hex()+"#"+comment.ID.Hex())
}

func ReplyComment(c echo.Context) error {
	//id := c.Param("id")

	return c.Redirect(http.StatusFound, "/article/")
}
