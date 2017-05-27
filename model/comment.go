package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Comment struct {
		ID        bson.ObjectId          `json:"id" bson:"_id,omitempty" xml:"id" form:"id" query:"name"`
		Editor    bson.ObjectId          `json:"editor" bson:"editor" xml:"editor" form:"editor" query:"editor"`
		Article   bson.ObjectId          `json:"article" bson:"article" xml:"article" form:"article" query:"article"`
		Replyto   bson.ObjectId          `json:"replyto" bson:"replyto"`
		Content   string                 `json:"content" bson:"content" xml:"content" form:"content" query:"content"`
		Like      int                    `json:"like" bson:"like" xml:"like" form:"like" query:"like"`
		UserLiked map[bson.ObjectId]bool `json:"user_liked" bson:"user_liked"`
		Replies   []bson.ObjectId        `json:"replies" bson:"replies" xml:"replies" form:"replies" query:"replies"`
		Time      time.Time              `json:"time" bson:"time" xml:"time" form:"time" query:"time"`
	}
)
