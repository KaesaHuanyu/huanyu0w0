package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Post struct {
		ID      bson.ObjectId `bson:"_id,omitempty"`
		Time    time.Time     `json:"time,omitempty" bson:"time,omitempty"`
		To      string        `json:"to" bson:"to"`
		From    string        `json:"from" bson:"from"`
		Message string        `json:"message" bson:"message"`
	}
)
