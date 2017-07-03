package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Log struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Object    string        `json:"object" bson:"object"`
		Type      string        `json:"type" bson:"type"`
		User      string        `json:"user" bson:"user"`
		Operation string        `json:"operation" bson:"operation"`
		Time      time.Time     `json:"time" bson:"time"`
		Signup    bool          `json:"signup" bson:"signup"`
	}
)

//[Time] [Admin] User Operation [Type] Object
