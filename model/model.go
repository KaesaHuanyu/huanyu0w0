package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"time"
)

const (
	MONGO_ADDRESS = "120.24.253.180:2333"
	MONGO_DB      = "huanyu0w0"
	MONGO_USER    = "user"
	MONGO_ARTICLE = "article"
	MONGO_COMMENT = "comment"
	COOKIE_TIME   = 24 * time.Hour
	REMEMBER      = 30 * 24 * time.Hour
)

var GlobalMgoSession *mgo.Session

////初始化连接池
//func init() {
//	globalMgoSession, err := mgo.Dial(MONGO_ADDRESS)
//	if err != nil {
//		panic(err)
//	}
//	GlobalMgoSession = globalMgoSession
//	GlobalMgoSession.SetMode(mgo.Monotonic, true)
//	//设置最大连接数
//	GlobalMgoSession.SetPoolLimit(2048)
//}
//
//获得连接
func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

type Cookies struct {
	IsLogin  bool
	UserId   string
	Avatar   string
	UserName string
}

type CommentEditor struct {
	Comment     *Comment
	Reply       *Comment
	Editor      *User
	Time        string
	ReplyNumber int
	Nice        bool
}

type ArticleEditor struct {
	Article      *Article
	Editor       *User
	Introduction template.HTML
}

func Insert(collection string, data interface{}) error {
	session := CloneSession()
	defer session.Close()

	c := session.DB(MONGO_DB).C(collection)
	err := c.Insert(data)
	if err != nil {
		log.Println("Insert data error.")
		return err
	}
	log.Println("insert data to Mongo successed!")
	return nil
}

func FindOne(collection string, id bson.ObjectId, data interface{}) error {
	session := CloneSession()
	defer session.Close()

	c := session.DB(MONGO_DB).C(collection)
	err := c.FindId(id).One(data)
	if err != nil {
		log.Println("Find data error.")
		return err
	}
	log.Println("find data from mongo successed")
	return nil
}

func FindOneNoId(collection string, key string, value string, data interface{}) error {
	session := CloneSession()
	defer session.Close()

	c := session.DB(MONGO_DB).C(collection)
	err := c.Find(bson.M{key: value}).One(data)
	if err != nil {
		log.Println("Find data error.")
		return err
	}
	log.Println("find data from mongo successed")
	return nil
}

func Update(collection string, id bson.ObjectId, data interface{}) error {
	session := CloneSession()
	defer session.Close()

	c := session.DB(MONGO_DB).C(collection)
	err := c.UpdateId(id, data)
	if err != nil {
		log.Println("Update data error.")
		return err
	}
	log.Println("update data from mongo successed")
	return nil
}

func Remove(collection string, id bson.ObjectId) error {
	session := CloneSession()
	defer session.Close()

	c := session.DB(MONGO_DB).C(collection)
	err := c.RemoveId(id)
	if err != nil {
		log.Println("Remove data error.")
		return err
	}
	log.Println("remove data from mongo successed")
	return nil
}

func FindAll(collection string, key string, value string, n int, sort string, data *[]*Article) error {
	session := CloneSession()
	defer session.Close()

	c := session.DB(MONGO_DB).C(collection)
	if key != "" {
		//查找排序前n条
		err := c.Find(bson.M{key: value}).Sort(sort).Limit(n).All(data)
		if err != nil {
			log.Println("Find data error.")
			return err
		}
	} else {
		//查找排序前n条
		err := c.Find(nil).Sort(sort).Limit(n).All(data)
		if err != nil {
			log.Println("Find data error.")
			return err
		}
	}
	log.Println("find data from mongo successed")
	return nil
}
