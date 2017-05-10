package model

import (
	"time"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
	"gopkg.in/mgo.v2/bson"
	"html/template"
)

const (
	MONGO_ADDRESS = "120.24.253.180:27018"
	MONGO_DB = "huanyu0w0"
	MONGO_USER = "user"
	MONGO_ARTICLE = "article"
	MONGO_COMMENT = "comment"
	COOKIE_TIME = 24 * time.Hour
	REMEMBER = 30 * 24 * time.Hour
)

type Article struct {
	Id string `json:"id" bson:"_id" xml:"id" form:"id" query:"id"`
	Time time.Time `json:"time" bson:"time" xml:"time" form:"time" query:"time"`
	Change time.Time `json:"change" bson:"change" xml:"change" form:"change" query:"change"`
	Editor string `json:"editor" bson:"editor" xml:"editor" form:"editor" query:"editor"`
	Topic string `json:"topic" bson:"topic" xml:"topic" form:"topic" query:"topic"`
	Title string `json:"title" bson:"title" xml:"title" form:"title" query:"title"`
	Introduction string `json:"introduction" bson:"introduction" xml:"introduction" form:"introduction" query:"introduction"`
	Content string `json:"content" bson:"content" xml:"content" form:"content" query:"content"`
	Comments []string `json:"comments" bson:"comments" xml:"comments" form:"comments" query:"comments"`
	Like int `json:"like" bson:"like" xml:"like" form:"like" query:"like"`
	UserLiked map[string]bool `json:"user_liked" bson:"user_liked"`
}

type User struct {
	Id string `json:"id" bson:"_id" xml:"id" form:"id" query:"name"`
	Time time.Time `json:"time" bson:"time" xml:"time" form:"time" query:"time"`
	Change time.Time `json:"change" bson:"change" xml:"change" form:"change" query:"change"`
	Email string `json:"email" bson:"email" xml:"email" form:"email" query:"email"`
	Name string `json:"name" bson:"name" xml:"name" form:"name" query:"name"`
	Password string `json:"password" bson:"password" xml:"password" form:"password" query:"name"`
	Avatar string `json:"avatar" bson:"avatar" xml:"avatar" form:"avatar" query:"avatar"`
	Comments []string `json:"comments" bson:"comments" xml:"comments" form:"comments" query:"comments"`
	Articles []string `json:"articles" bson:"articles" xml:"articles" form:"articles" query:"articles"`
}

type Comment struct {
	Id string `json:"id" bson:"_id" xml:"id" form:"id" query:"name"`
	Editor string `json:"editor" bson:"editor" xml:"editor" form:"editor" query:"editor"`
	Article string `json:"article" bson:"article" xml:"article" form:"article" query:"article"`
	Content string `json:"content" bson:"content" xml:"content" form:"content" query:"content"`
	Like int `json:"like" bson:"like" xml:"like" form:"like" query:"like"`
	UserLiked map[string]bool `json:"user_liked" bson:"user_liked"`
	Comment string `json:"comment" bson:"comment" xml:"comment"`
	Replies []string `json:"replies" bson:"replies" xml:"replies" form:"replies" query:"replies"`
	Time time.Time `json:"time" bson:"time" xml:"time" form:"time" query:"time"`
}

type Image struct {
	Id string `json:"id" xml:"id" bson:"_id" form:"id" query:"name"`
	Url string `json:"url" xml:"url" bson:"url"`
}

type Cookies struct {
	IsLogin bool
	UserId string
	Avatar string
	UserName string
}

type CommentEditor struct {
	Comment *Comment
	Reply *Comment
	Editor *User
	Time string
	ReplyNumber int
	Nice bool
}

type ArticleEditor struct {
	Article *Article
	Editor *User
	Introduction template.HTML
}

func InsertMongo(collection string, data interface{}) error {
	session, err := mgo.Dial(MONGO_ADDRESS)
	if session == nil {
		return fmt.Errorf("error: session is nil")
	}
	if err != nil {
		log.Println("Dial mongo error.")
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MONGO_DB).C(collection)
	err = c.Insert(data)
	if err != nil {
		log.Println("Insert data error.")
		return err
	}
	log.Println("insert data to Mongo successed!")
	return nil
}

func FindMongo(collection string, key string, value string, data interface{}) error {
	session, err := mgo.Dial(MONGO_ADDRESS)
	if session == nil {
		return fmt.Errorf("error: session is nil")
	}
	if err != nil {
		log.Println("Dial mongo error.")
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(MONGO_DB).C(collection)
	err = c.Find(bson.M{key:value}).One(data)
	if err != nil {
		log.Println("Find data error.")
		return err
	}
	log.Println("find data from mongo successed")
	return nil
}

func UpdateMongo(collection string, id string, data interface{}) error {
	session, err := mgo.Dial(MONGO_ADDRESS)
	if session == nil {
		return fmt.Errorf("error: session is nil")
	}
	if err != nil{
		log.Println("Dial mongo error.")
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(MONGO_DB).C(collection)
	err = c.UpdateId(id, data)
	if err != nil {
		log.Println("Update data error.")
		return err
	}
	log.Println("update data from mongo successed")
	return nil
}

func RemoveMongo(collection string, id string) error {
	session, err := mgo.Dial(MONGO_ADDRESS)
	if session == nil {
		return fmt.Errorf("error: session is nil")
	}
	if err != nil {
		log.Println("Dial mongo error.")
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MONGO_DB).C(collection)
	err = c.RemoveId(id)
	if err != nil {
		log.Println("Remove data error.")
		return err
	}
	log.Println("remove data from mongo successed")
	return nil
}

func FindAll(collection string, key string, value string, n int, sort string, data *[]*Article) error {
	session, err := mgo.Dial(MONGO_ADDRESS)
	if session == nil {
		return fmt.Errorf("error: session is nil")
	}
	if err != nil {
		log.Println("Dial mongo error.")
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MONGO_DB).C(collection)
	if key != "" {
		//查找排序前n条
		err = c.Find(bson.M{key:value}).Sort(sort).Limit(n).All(data)
	} else {
		//查找排序前n条
		err = c.Find(nil).Sort(sort).Limit(n).All(data)
	}
	if err != nil {
		log.Println("Find data error.")
		return err
	}
	log.Println("find data from mongo successed")
	return nil
}

//func ReadCookie(c echo.Context, cookiename string) error {
//	cookie, err := c.Cookie(cookiename)
//	if err != nil {
//		return err
//	}
//
//
//}