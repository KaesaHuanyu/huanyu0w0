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
	KEY = "secret"
	MONGO_ADDRESS = "120.24.253.180:2333"
	MONGO_DB = "huanyu0w0"
	USER = "users"
	ARTICLE = "articles"
	COMMENT = "comments"
	POST = "posts"
)