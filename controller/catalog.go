package controller

import (
	"github.com/labstack/echo"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
)

func Catalog(c echo.Context) error {
	//return c.HTML(http.StatusOK, "<!DOCTYPE html><html> <head> <title>寰宇0_0 - 哇咔咔</title> </head> <body> <img src=\"../static/img/二娃.jpg\"> <h1>大家好，我是二娃</h1> </body> </html>")
	input, err := ioutil.ReadFile("../static/test.md")
	if err != nil {
		log.Println(err)
	}
	output := blackfriday.MarkdownBasic(input)
	log.Println(string(output))
	return c.Render(http.StatusOK, "catalog", string(output)) //name在模版中定义
}
