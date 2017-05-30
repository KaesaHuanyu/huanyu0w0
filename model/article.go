package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Article struct {
		ID        bson.ObjectId   `json:"id" bson:"_id,omitempty"`
		Time      time.Time       `json:"time" bson:"time"`
		Change    time.Time       `json:"change,omitempty" bson:"change,omitempty"`
		Editor    string          `json:"editor" bson:"editor"`
		Topic     string          `json:"topic" bson:"topic" form:"topic"`
		Title     string          `json:"title" bson:"title" form:"title"`
		Reason    string          `json:"reason" bson:"reason" form:"reason"`
		Comments  []string        `json:"comments" bson:"comments"`
		Like      int             `json:"like" bson:"like"`
		UserLiked map[string]bool `json:"user_liked" bson:"user_liked"`
	}

	Display struct {
		Article *Article
		Editor *User
	}
)
