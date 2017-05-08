package main

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	//"huanyu0w0/model"
	//"fmt"
	//"time"
	"huanyu0w0/model"
	//"fmt"
	"time"
	"os"
)

func main() {
	session, err := mgo.Dial("120.24.253.180:27017")
	if err != nil {
		panic(err)
	}
	if session == nil {
		panic("session is nil!")
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("testDB").C("testC")
	//insert
	err = c.Insert(&model.User{
		//Id: os.Geteuid(),
		Time: time.Now(),
		Name: "huanyu0w0",
		Password: "3.1415926",
	})
	if err != nil {
		log.Println("insert error")
		log.Fatal(err)
	}

	//find
	//result := model.User{}
	//change := mgo.Change{
	//	Update:bson.M{"$inc":bson.M{"_id":1}},
	//	ReturnNew:true,
	//}
	//_, err = c.Find(bson.M{"_id": 1}).Apply(change, &result)
	//if err != nil {
	//	log.Println("find error")
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(result)
	//
	//err = c.Remove(bson.M{"name":"huanyu0w0"})
	//if err != nil {
	//	log.Println("here")
	//	log.Println(err)
	//}
	//
	//err = c.Find(bson.M{"name":"huanyu0w0"}).One(&result)
	//if err != nil {
	//	log.Println("find error")
	//	log.Fatal(err)
	//}
}