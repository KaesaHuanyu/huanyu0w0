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
	MONGO_ADDRESS = "huanyu0w0-mongo:27017"
	MONGO_DB      = "huanyu0w0"
	USER          = "users"
	ARTICLE       = "articles"
	COMMENT       = "comments"
	LOG           = "logs"
)
