package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		ShowID string `json:"show_id,omitempty" bson:"show_id,omitempty"`
		Time      time.Time     `json:"time" bson:"time"`
		Change    time.Time     `json:"change,omitempty" bson:"change,omitempty"`
		Email     string        `json:"email" bson:"email" form:"email"`
		Name      string        `json:"name" bson:"name" form:"name"`
		Password  string        `json:"password,omitempty" bson:"password,omitempty" form:"password"`
		Avatar    string        `json:"avatar" bson:"avatar" form:"avatar"`
		BigAvatar string `json:"big_avatar,omitempty" bson:"big_avatar,omitempty"`
		BGM string `json:"bgm,omitempty" bson:"bgm,omitempty" form:"bgm"`
		Follows   []string      `json:"follow" bson:"follow"`
		Comments  []string      `json:"comments" bson:"comments"`
		Articles  []string      `json:"articles" bson:"articles"`
		History []string `json:"history" bson:"history"`
		Collection []string `json:"collection" bson:"collection"`
		Follower int      `json:"follower" bson:"follower"`
		IsFollower map[string]bool `json:"is_follower" bson:"is_follower"`
	}

	UserDisplay struct {
		User      *User
		ID string
		CreateTime string
		Articles  []*Article
		Comments  []*Comment
		Follow    []*User
		Followers []*User
	}

	Admin struct {
		User
		AdminPassword string
	}
)

func (u *User) GetCreateTime() (createTime string) {
	return u.Time.Format("2006年 01月02日 15:04")
}
