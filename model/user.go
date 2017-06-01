package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		//Token     string        `json:"token,omitempty" bson:"-"`
		Time      time.Time     `json:"time" bson:"time"`
		Change    time.Time     `json:"change,omitempty" bson:"change,omitempty"`
		Email     string        `json:"email" bson:"email" form:"email"`
		Name      string        `json:"name" bson:"name" form:"name"`
		Password  string        `json:"password,omitempty" bson:"password,omitempty" form:"password"`
		Avatar    string        `json:"avatar" bson:"avatar" form:"avatar"`
		Follows   []string      `json:"follow" bson:"follow"`
		Comments  []string      `json:"comments" bson:"comments"`
		Articles  []string      `json:"articles" bson:"articles"`
		Followers []string      `json:"followers" bson:"followers"`
	}
)