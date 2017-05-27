package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID    bson.ObjectId `json:"id" bson:"_id,omitempty" xml:"id" form:"id" query:"name"`
		Token string        `json:"token,omitempty" bson:"-"`
		//Time time.Time `json:"time" bson:"time" xml:"time" form:"time" query:"time"`
		//Change time.Time `json:"change" bson:"change" xml:"change" form:"change" query:"change"`
		Email     string          `json:"email" bson:"email" xml:"email" form:"email" query:"email"`
		Name      string          `json:"name" bson:"name" xml:"name" form:"name" query:"name"`
		Password  string          `json:"password" bson:"password" xml:"password" form:"password" query:"name"`
		Avatar    string          `json:"avatar" bson:"avatar" xml:"avatar" form:"avatar" query:"avatar"`
		Comments  []bson.ObjectId `json:"comments" bson:"comments" xml:"comments" form:"comments" query:"comments"`
		Articles  []bson.ObjectId `json:"articles" bson:"articles" xml:"articles" form:"articles" query:"articles"`
		Followers []string        `json:"followers,omitempty" bson:"followers,omitempty"`
	}
)
