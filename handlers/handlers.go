package handlers

import (
	"gopkg.in/mgo.v2"
)

type (
	Handler struct {
		DB *mgo.Session
	}
)

const (
	MONGO_ADDRESS = "120.24.253.180:32786"
	MONGO_DB      = "huanyu0w0"
	USER          = "users"
	ARTICLE       = "articles"
	COMMENT       = "comments"
	POST          = "posts"
)

var (
	BGM = []string{
		"30245064",
	}
	WallPaper = []string{}
)
