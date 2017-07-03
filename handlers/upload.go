package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"io"
	"net/http"
	"os"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"strings"
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
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "请先选择图片"}
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println()
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}
	defer src.Close()
	//Destination
	filePath := "static/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println("os.Create:")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("io.Copy")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	originAvatar := strings.TrimPrefix(strings.TrimSuffix(data.Avatar, "-avatarStyle"),
		"http://images.huanyu0w0.cn/")
	avatar, err := toQiniu(data.ID, file.Filename, filePath, originAvatar, "-avatarStyle")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("os.Remove:")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$set": bson.M{"avatar": avatar}}); err != nil {
		if err == mgo.ErrNotFound {
			fmt.Println("UpdateId:")
			return echo.ErrNotFound
		} else {
			fmt.Println("UpdateId:")
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	ck, _ := c.Cookie("avatar")
	ck.Value = avatar
	c.SetCookie(ck)

	return c.Redirect(http.StatusFound, "/user/dashboard")
}

func toQiniu(id, fileName, filePath, originAvatar, style string) (url string, err error) {
	conf.ACCESS_KEY = model.ACCESS_KEY
	conf.SECRET_KEY = model.SECRET_KEY
	key := "avatar/" + id + "/" + fileName

	// 创建一个Client
	c := kodo.New(0, nil)
	//删除原文件
	p := c.Bucket(model.BUCKET)
	// 调用Delete方法删除文件
	err = p.Delete(nil, originAvatar)
	// 打印返回值以及出错信息
	if err == nil {
		fmt.Println("Delete success")
	} else {
		fmt.Println(err)
	}
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: model.BUCKET + ":" + key,
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
	url = "http://images.huanyu0w0.cn/" + key + style
	return
}
