package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type User struct {
	ID bson.ObjectId `bson:"_id"`
	Name string
	Articles []bson.ObjectId
}

type Article struct {
	ID bson.ObjectId `bson:"_id"`
	Title string
	Content string
	Editor bson.ObjectId
}

var GlobalMgoSession *mgo.Session

func init() {
	globalMgoSession, err := mgo.Dial("120.24.253.180:23333")
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//设置最大连接数
	GlobalMgoSession.SetPoolLimit(300)
}

func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

func main() {
	session := CloneSession()
	defer session.Close()

	c1 := session.DB("huanyu0w0").C("User")
	c2 := session.DB("huanyu0w0").C("Article")
	//u := &User{
	//	ID: bson.NewObjectId(),
	//	Name: "huanyu0w0",
	//}
	//a := &Article{
	//	ID: bson.NewObjectId(),
	//	Title: "hello",
	//	Content: "world",
	//	Editor: u.ID,
	//}
	//u.Articles = append(u.Articles, a.ID)
	//if err := c1.Insert(u); err != nil {
	//	log.Fatal(err)
	//}
	//if err := c2.Insert(a); err != nil {
	//	log.Fatal(err)
	//}
	u := &User{}
	c1.Find(bson.M{"name":"huanyu0w0"}).One(u)
	fmt.Println(u)
	a := &Article{}
	c2.FindId(u.Articles[0]).One(a)
	fmt.Println(a)
}