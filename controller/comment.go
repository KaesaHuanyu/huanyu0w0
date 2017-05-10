package controller

import (
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"net/http"
	"time"
	"github.com/satori/go.uuid"
	"log"
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
		Id: uuid.NewV4().String(),
		Time: time.Now(),
		Article: article,
		Editor: editor.Value,
		Content: content,
	}

	if err := c.Bind(comment); err != nil {
		log.Println("CreateComment Bind error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	err = model.InsertMongo(model.MONGO_COMMENT, comment)
	if err != nil {
		log.Println("CreateComment InsertMongo error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}

	a := new(model.Article)
	err = model.FindMongo(model.MONGO_ARTICLE, "_id", article, a)
	if err != nil {
		log.Println("CreateComment FindMongo article error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	a.Comments = append(a.Comments, comment.Id)
	err = model.UpdateMongo(model.MONGO_ARTICLE, article, a)
	if err != nil {
		log.Println("CreateComment UpdateMongo article error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}

	u := new(model.User)
	err = model.FindMongo(model.MONGO_USER, "_id", editor.Value, u)
	if err != nil {
		log.Println("CreateComment FindMongo editor error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	u.Comments = append(u.Comments, comment.Id)
	err = model.UpdateMongo(model.MONGO_USER, editor.Value, u)
	if err != nil {
		log.Println("CreateComment UpdateMongo editor error.", err)
		return c.Render(http.StatusFound, "error", "评论失败...")
	}
	return c.Redirect(http.StatusFound, "/article/" + comment.Article)
}

func GetComment(c echo.Context) error {
	id := c.Param("id")
	comment := new(model.Comment)
	err := model.FindMongo(model.MONGO_COMMENT, "_id", id, comment)
	if err != nil {
		log.Println("GetComment FindMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

func DeleteComment(c echo.Context) error {
	id := c.Param("id")
	err := model.RemoveMongo(model.MONGO_COMMENT, id)
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
	err = model.FindMongo(model.MONGO_COMMENT, "_id", id, comment)
	if err != nil {
		log.Println("NiceComment FindMongo error, ", err)
		return c.Render(http.StatusFound, "error", "赞失败了...找不到这条评论...")
	}
	if comment.UserLiked[user.Value] == true {
		comment.Like--
		comment.UserLiked[user.Value] = false
		nice.Value = "false"
		c.SetCookie(nice)
		err := model.UpdateMongo(model.MONGO_COMMENT, id, comment)
		if err != nil {
			log.Println("NiceComment Update error: ", err)
			c.Render(http.StatusFound, "error", "牙白...出了点问题...")
		}
		return c.Redirect(http.StatusFound, "/article/" + comment.Article + "#" + comment.Id)
	}
	comment.Like++
	comment.UserLiked[user.Value] = true
	nice.Value = "true"
	err = model.UpdateMongo(model.MONGO_COMMENT, id, comment)
	if err != nil {
		log.Println("NiceComment Update error: ", err)
		c.Render(http.StatusFound, "error", "牙白...出了点问题...")
	}
	return c.Redirect(http.StatusFound, "/article/" + comment.Article + "#" + comment.Id)
}

func ReplyComment(c echo.Context) error {
	//id := c.Param("id")

	return c.Redirect(http.StatusFound, "/article/")
}