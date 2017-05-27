package controllers

import (
	"github.com/labstack/echo"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"huanyu0w0/model"
	"log"
	"net/http"
	"strconv"
	"time"
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
	err = model.FindOne(model.MONGO_USER, bson.ObjectIdHex(userid.Value), u)
	if err != nil {
		return c.Render(http.StatusFound, "error", err)
	}

	a := &model.Article{
		ID:     bson.NewObjectId(),
		Time:   time.Now(),
		Editor: u.ID,
	}
	if err := c.Bind(a); err != nil {
		log.Println("CreateArticle context.Bind() error.")
		return c.Render(http.StatusFound, "error", err)
	}
	//User中插入article
	u.Articles = append(u.Articles, a.ID)
	err = model.Update(model.MONGO_USER, u.ID, u)
	if err != nil {
		log.Println("CreateArticle model.UpdateMongo() error.", err)
		return c.Render(http.StatusFound, "error", "保存安利💊失败...")
	}
	//插入article到数据库
	err = model.Insert(model.MONGO_ARTICLE, a)
	if err != nil {
		log.Println("CreateArticle model.InsertMongo() error.", err)
		return c.Render(http.StatusFound, "error", "保存安利💊失败...")
	}
	log.Println(a)
	return c.Redirect(http.StatusFound, "/article/"+a.ID.Hex())
}

func GetArticle(c echo.Context) error {
	data := &struct {
		model.Cookies
		Article  *model.Article
		Editor   *model.User
		Time     string
		Comments []*model.CommentEditor
		Content  template.HTML
		Like     bool
	}{
		Article:  &model.Article{},
		Editor:   &model.User{},
		Comments: []*model.CommentEditor{},
	}
	//处理cookie
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
		like, err := c.Cookie("like")
		if err == nil {
			if like.Value == "true" {
				data.Like = true
			}
		}
	}

	time := time.Now()
	id := c.Param("id")
	err = model.FindOne(model.MONGO_ARTICLE, bson.ObjectIdHex(id), data.Article)
	if err != nil {
		log.Println("GetArticle FindMongo Article error: ", err)
		return c.Render(http.StatusFound, "error", "404 NOT FOUND")
	}
	if data.Article.Time.Year() == time.Year() {
		if data.Article.Time.Month() == time.Month() && data.Article.Time.Day() == time.Day() {
			if data.Article.Time.Hour() == time.Hour() {
				if data.Article.Time.Minute() == time.Minute() {
					data.Time = strconv.Itoa(time.Second()-data.Article.Time.Second()) + "秒前"
				} else {
					data.Time = strconv.Itoa(time.Minute()-data.Article.Time.Minute()) + "分钟前"
				}
			} else {
				data.Time = data.Article.Time.Format("15:04")
			}
		} else {
			data.Time = data.Article.Time.Format("01月02日 15:04")
		}
	} else {
		data.Time = data.Article.Time.Format("2006年 01月02日 15:04")
	}
	data.Content = template.HTML(blackfriday.MarkdownCommon([]byte(data.Article.Content)))

	err = model.FindOne(model.MONGO_USER, data.Article.Editor, data.Editor)
	if err != nil {
		log.Println("GetArticle FindMongo User error: ", err)
		return c.Render(http.StatusFound, "error", "404 NOT FOUND")
	}
	for index, comment := range data.Article.Comments {
		data.Comments = append(data.Comments, &model.CommentEditor{
			&model.Comment{},
			&model.Comment{},
			&model.User{},
			"",
			0,
			false,
		})
		err = model.FindOne(model.MONGO_COMMENT, comment, data.Comments[index].Comment)
		if err != nil {
			log.Println("GetArticle FindMongo CommentEditor.Comment error: ", err)
			return c.Render(http.StatusFound, "error", "404 NOT FOUND")
		}
		nice := new(http.Cookie)
		nice.Name = comment.Hex()

		//实现装逼的时间显示效果
		if data.Comments[index].Comment.Time.Year() == time.Year() {
			if data.Comments[index].Comment.Time.Month() == time.Month() && data.Comments[index].Comment.Time.Day() == time.Day() {
				if data.Comments[index].Comment.Time.Hour() == time.Hour() {
					if data.Comments[index].Comment.Time.Minute() == time.Minute() {
						data.Comments[index].Time = strconv.Itoa(time.Second()-data.Comments[index].Comment.Time.Second()) + "秒前"
					} else {
						data.Comments[index].Time = strconv.Itoa(time.Minute()-data.Comments[index].Comment.Time.Minute()) + "分钟前"
					}
				} else {
					data.Comments[index].Time = data.Comments[index].Comment.Time.Format("15:04")
				}
			} else {
				data.Comments[index].Time = data.Comments[index].Comment.Time.Format("01月02日 15:04")
			}
		} else {
			data.Comments[index].Time = data.Comments[index].Comment.Time.Format("2006年 01月02日 15:04")
		}

		//获取回复数
		data.Comments[index].ReplyNumber = len(data.Comments[index].Comment.Replies)

		err = model.FindOne(model.MONGO_USER, data.Comments[index].Comment.Editor, data.Comments[index].Editor)
		if err != nil {
			log.Println("GetArticle FindMongo CommentEditor.Editor error: ", err)
			return c.Render(http.StatusFound, "error", "404 NOT FOUND")
		}
		if data.Comments[index].Comment.UserLiked[bson.ObjectIdHex(data.UserId)] == true {
			data.Comments[index].Nice = true
		}
	}
	if data.Article.UserLiked[bson.ObjectIdHex(data.UserId)] == true {
		data.Like = true
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
	err := model.Update(model.MONGO_ARTICLE, bson.ObjectIdHex(id), a)
	if err != nil {
		log.Println("UpdateArticle UpdateMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func DeleteArticle(c echo.Context) error {
	id := c.Param("id")
	err := model.Remove(model.MONGO_ARTICLE, bson.ObjectIdHex(id))
	if err != nil {
		log.Println("DeleteArticle RemoveMongo error.")
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func LikeArticle(c echo.Context) error {
	userid, err := c.Cookie("userid")
	if err != nil {
		log.Println("not sign in, ERROR: ", err)
		return c.Redirect(http.StatusFound, "/signin")
	}
	like, _ := c.Cookie("like")

	id := c.Param("id")
	a := new(model.Article)
	err = model.FindOne(model.MONGO_ARTICLE, bson.ObjectIdHex(id), a)
	if err != nil {
		log.Println("GetArticle FindMongo error.")
		return c.Render(http.StatusNotFound, "error", "没找到...")
	}
	if a.UserLiked[bson.ObjectIdHex(userid.Value)] == true {
		a.Like--
		a.UserLiked[bson.ObjectIdHex(userid.Value)] = false
		like.Value = "false"
		c.SetCookie(like)
		err = model.Update(model.MONGO_ARTICLE, bson.ObjectIdHex(id), a)
		if err != nil {
			log.Println("LikeArticle Update error: ", err)
			c.Render(http.StatusFound, "error", "牙白...出了点问题...")
		}
		return c.Redirect(http.StatusFound, "/article/"+id)
	}
	a.Like++
	a.UserLiked[bson.ObjectIdHex(userid.Value)] = true
	like.Value = "true"
	c.SetCookie(like)
	err = model.Update(model.MONGO_ARTICLE, bson.ObjectIdHex(id), a)
	if err != nil {
		log.Println("Update error: ", err)
		c.Render(http.StatusFound, "error", "牙白...出了点问题...")
	}
	return c.Redirect(http.StatusFound, "/article/"+id)
}
