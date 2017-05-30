package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Comment struct {
		ID        bson.ObjectId   `json:"id" bson:"_id,omitempty"`
		Editor    string          `json:"editor" bson:"editor"`
		Article   string          `json:"article" bson:"article"`
		Replyto   string          `json:"replyto" bson:"replyto"`
		Content   string          `json:"content" bson:"content" form:"content"`
		Like      int             `json:"like" bson:"like"`
		UserLiked map[string]bool `json:"user_liked" bson:"user_liked"`
		Replies   []string        `json:"replies" bson:"replies"`
		Time      time.Time       `json:"time" bson:"time"`
	}
)
