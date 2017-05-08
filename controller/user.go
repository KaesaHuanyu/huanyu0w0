package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"huanyu0w0/model"
	"time"
	"github.com/satori/go.uuid"
	"log"
	"html/template"
)

//创建用户
func CreateUser(c echo.Context) error {
	must := c.FormValue("2333")
	if must != "on" {
		return c.Render(http.StatusFound, "error", "请勾选誓约")
	}
	u := &model.User{
		Id: uuid.NewV4().String(),//生成uuid
		Time: time.Now(),//当前服务器时间为创建时间
		Avatar: "http://images.huanyu0w0.cn/icon.jpg",
	}
	if err := c.Bind(u); err != nil {
		log.Println("CreateUser Bind error.")
		return c.Redirect(http.StatusFound, "/signup")
	}
	//重复验证
	if err := model.FindMongo(model.MONGO_USER, "email", u.Email, nil); err == nil {
		return c.Render(http.StatusFound, "error", template.HTML("邮箱已存在"))
	}
	if err := model.FindMongo(model.MONGO_USER, "name", u.Name, nil); err == nil {
		return c.Render(http.StatusFound, "error", template.HTML("用户名已存在"))
	}
	//插入数据
	err := model.InsertMongo(model.MONGO_USER, u)
	if err != nil {
		log.Println("CreateUser InsertMongo error.")
		return c.Redirect(http.StatusFound, "/signup")
	}
	log.Println(u)
	//return c.JSON(http.StatusOK, u)
	return c.Redirect(http.StatusOK, "/")
}

//查找用户
func GetUser(c echo.Context) error {
	id := c.Param("id")
	u := new(model.User)
	err := model.FindMongo(model.MONGO_USER, "_id", id, u)
	if err != nil {
		log.Println("GetUser FindMongo error.")
		return c.NoContent(http.StatusNotFound)
	}
	//return c.JSON(http.StatusOK, u)
	data := struct {
		user *model.User
		isSignin bool
	}{
		user: u,
		isSignin: false,
	}
	_, err = c.Cookie("isSignin")
	if err != nil {
		log.Println("not sign in, ERROR: ", err)
	} else {
		data.isSignin = true
	}
	return c.Render(http.StatusOK, "user", data)
}

//更新用户信息
func UpdateUser(c echo.Context) error {
	u := &model.User{
		Change: time.Now(),
	}
	if err := c.Bind(u); err != nil {
		log.Println("UpdateUser Bind error.")
		return err
	}
	id := c.Param("id")
	err := model.UpdateMongo(model.MONGO_USER, id, u)
	if err != nil {
		log.Println("UpdateUser UpdateMongo error.")
		return err
	}
	return c.JSON(http.StatusOK, u)
}

//删除用户
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := model.RemoveMongo(model.MONGO_USER, id)
	if err != nil {
		log.Println("DeleteUser RemoveMongo error.")
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
