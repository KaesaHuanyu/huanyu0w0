package handlers

import (
	"github.com/labstack/echo"
	"os"
	"io"
	"fmt"
	"net/http"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"huanyu0w0/model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func (h *Handler) UpdateAvatar(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}
	//Read file
	//Source
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("FormFile:")
		return
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println()
		return
	}
	defer src.Close()
	//Destination
	filePath := "static/tmp/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println("os.Create:")
		return
	}

	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("io.Copy")
		return
	}

	avatar, err := toQiniu(data.ID, file.Filename, filePath)
	if err != nil {
		fmt.Println()
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("os.Remove:")
		return
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB(MONGO_DB).C(USER).
	UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$set": bson.M{"avatar": avatar}}); err != nil{
		if err == mgo.ErrNotFound {
			fmt.Println("UpdateId:")
			return echo.ErrNotFound
		} else {
			fmt.Println("UpdateId:")
			return err
		}
	}

	ck, err := c.Cookie("avatar")
	if err != nil {
		return
	}
	ck.Value = avatar
	c.SetCookie(ck)

	return c.Redirect(http.StatusFound, "/")
}

func toQiniu(id, fileName, filePath string) (url string, err error) {
	conf.ACCESS_KEY = model.ACCESS_KEY
	conf.SECRET_KEY = model.SECRET_KEY
	key := "avatar/" + id + "/" + fileName

	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope:   model.BUCKET+ ":" + key,
		//设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, &kodocli.UploadConfig{
		UpHosts: []string{"http://up-z2.qiniu.com", "http://up.qiniu.com"},
	})

	var ret model.PutRet
	// 调用PutFile方式上传，这里的key需要和上传指定的key一致
	err = uploader.PutFile(nil, &ret, token, key, filePath, nil)
	if err != nil {
		fmt.Println("uploader.PutFile:")
		return
	}
	url = "http://images.huanyu0w0.cn/" + key + "-avatarStyle"
	return
}