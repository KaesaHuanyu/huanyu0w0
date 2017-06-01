package model

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
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
		Name      string          `json:"name" bson:"name" form:"name"`
		Reason    string          `json:"reason" bson:"reason" form:"reason"`
		Url       string          `json:"url,omitempty" bson:"url,omitempty" form:"url"`
		//Link      template.HTML   `json:"link,omitempty" bson:"link,omitempty" form:"link"`
		Comments  []string        `json:"comments" bson:"comments"`
		Like      int             `json:"like" bson:"like"`
		UserLiked map[string]bool `json:"user_liked" bson:"user_liked"`
	}

	Display struct {
		Article     *Article
		Editor      *User
		ID          string
		ShowTime    string
		CommentsNum int
	}
)

func (a *Article) GetShowTime() (showTime string) {
	time := time.Now()
	if a.Time.Year() == time.Year() {
		if a.Time.Month() == time.Month() && a.Time.Day() == time.Day() {
			if a.Time.Hour() == time.Hour() {
				if a.Time.Minute() == time.Minute() {
					showTime = strconv.Itoa(time.Second()-a.Time.Second()) + "秒前"
				} else {
					showTime = strconv.Itoa(time.Minute()-a.Time.Minute()) + "分钟前"
				}
			} else {
				showTime = "今天" + a.Time.Format(" 15:04")
			}
		} else if a.Time.YearDay() == time.YearDay() - 1 {
			showTime = "昨天" + a.Time.Format(" 15:04")
		} else {
			showTime = a.Time.Format("01月02日 15:04")
		}
	} else {
		showTime = a.Time.Format("2006年 01月02日 15:04")
	}
	return showTime
}
