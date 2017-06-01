package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
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

func (comment *Comment) GetShowTime() (showTime string) {
	time := time.Now()
	if comment.Time.Year() == time.Year() {
		if comment.Time.Month() == time.Month() && comment.Time.Day() == time.Day() {
			if comment.Time.Hour() == time.Hour() {
				if comment.Time.Minute() == time.Minute() {
					showTime = strconv.Itoa(time.Second() - comment.Time.Second()) + "秒前"
				} else {
					showTime = strconv.Itoa(time.Minute() - comment.Time.Minute()) + "分钟前"
				}
			} else {
				showTime = "今天" + comment.Time.Format("15:04")
			}
		} else if comment.Time.YearDay() == time.YearDay() - 1 {
			showTime = "昨天" + comment.Time.Format("15:04")
		} else {
			showTime = comment.Time.Format("01月02日 15:04")
		}
	} else {
		showTime = comment.Time.Format("2006年 01月02日 15:04")
	}

	return showTime
}