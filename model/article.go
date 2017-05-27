package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Article struct {
		ID        bson.ObjectId          `json:"id" bson:"_id,omitempty" xml:"id" form:"id" query:"id"`
		Time      time.Time              `json:"time" bson:"time" xml:"time" form:"time" query:"time"`
		Change    time.Time              `json:"change" bson:"change" xml:"change" form:"change" query:"change"`
		Editor    bson.ObjectId          `json:"editor" bson:"editor" xml:"editor" form:"editor" query:"editor"`
		Topic     string                 `json:"topic" bson:"topic" xml:"topic" form:"topic" query:"topic"`
		Title     string                 `json:"title" bson:"title" xml:"title" form:"title" query:"title"`
		Content   string                 `json:"content" bson:"content" xml:"content" form:"content" query:"content"`
		Comments  []bson.ObjectId        `json:"comments" bson:"comments" xml:"comments" form:"comments" query:"comments"`
		Like      int                    `json:"like" bson:"like" xml:"like" form:"like" query:"like"`
		UserLiked map[bson.ObjectId]bool `json:"user_liked" bson:"user_liked"`
	}
)
