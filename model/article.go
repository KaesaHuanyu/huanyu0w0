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
		Comments  []string        `json:"comments" bson:"comments"`
		Like      int             `json:"like" bson:"like"`
		UserLiked map[string]bool `json:"user_liked" bson:"user_liked"`
	}

	Display struct {
		Article     *Article
		Editor      *User
		Comments []*DisplayComment
		MostLikes *DisplayComment
		IsMostLikes bool
		ID          string
		ShowTime    string
		ShowTopic string
		CommentsNum int
		Fans int
		IsLike bool
	}
)

func (a *Article) GetShowTime() (showTime string) {
	time := time.Now()
	if a.Time.Year() == time.Year() {
		if a.Time.Month() == time.Month() && a.Time.Day() == time.Day() {
			if a.Time.Hour() == time.Hour() {
				if a.Time.Minute() == time.Minute() {
					if a.Time.Second() == time.Second() {
						showTime = "刚刚"
					} else {
						showTime = strconv.Itoa(time.Second()-a.Time.Second()) + "秒前"
					}
				} else if a.Time.Minute() == time.Minute()-1 && a.Time.Second() > time.Second() {
					showTime = strconv.Itoa(time.Second()+60-a.Time.Second()) + "秒前"
				} else {
					showTime = strconv.Itoa(time.Minute()-a.Time.Minute()) + "分钟前"
				}
			} else if a.Time.Hour() == time.Hour()-1 && a.Time.Minute() > time.Minute() {
				showTime = strconv.Itoa(time.Minute()+60-a.Time.Minute()) + "分钟前"
			} else {
				showTime = "今天" + a.Time.Format(" 15:04")
			}
		} else if a.Time.YearDay() == time.YearDay()-1 {
			showTime = "昨天" + a.Time.Format(" 15:04")
		} else {
			showTime = a.Time.Format("01月02日 15:04")
		}
	} else {
		showTime = a.Time.Format("2006年 01月02日 15:04")
	}
	return showTime
}

func (a *Article) GetShowTopic() (showTopic string) {
	switch a.Topic {
	case "anime":
		showTopic = "番剧"
	case "movie":
		showTopic = "电影"
	case "tv":
		showTopic = "影视"
	case "music_not":
		showTopic = "音乐"
	case "book":
		showTopic = "文字"
	case "videogame_asset":
		showTopic = "游戏"
	case "others":
		showTopic = "其他"
	}
	return showTopic
}
