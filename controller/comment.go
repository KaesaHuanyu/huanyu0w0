package controller

import (
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"net/http"
	"time"
	"github.com/satori/go.uuid"
	"log"
)

func CreateComment(c echo.Context) error {
	comment := &model.Comment{
		Id: uuid.NewV4().String(),
		Time: time.Now(),
	}
	if err := c.Bind(comment); err != nil {
		log.Println("CreateComment Bind error.")
		return err
	}
	err := model.InsertMongo(model.MONGO_COMMENT, comment)
	if err != nil {
		log.Println("CreateComment InsertMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

func GetComment(c echo.Context) error {
	id := c.Param("id")
	comment := new(model.Comment)
	err := model.FindMongo(model.MONGO_COMMENT, "_id", id, comment)
	if err != nil {
		log.Println("GetComment FindMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

func DeleteComment(c echo.Context) error {
	id := c.Param("id")
	err := model.RemoveMongo(model.MONGO_COMMENT, id)
	if err != nil {
		log.Println("DeleteComment RemoveMongo error.")
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
