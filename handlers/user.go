package handlers

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"net/http"
	"time"
)

func (h *Handler) SignupGet(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}
	return c.Render(http.StatusOK, "signup", data)
}

func (h *Handler) Signup(c echo.Context) (err error) {
	//清除cookies
	for _, cookie := range c.Cookies() {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}
	//Bind
	u := &model.User{
		ID:     bson.NewObjectId(),
		Time:   time.Now(),
		Avatar: "http://images.huanyu0w0.cn/icon.jpg",
	}

	if err = c.Bind(u); err != nil {
		return
	}
	//Validate
	if u.Email == "" || u.Password == "" || u.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "不能为空"}
	}

	//Check email
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB(MONGO_DB).C(USER).
		Find(bson.M{"email": u.Email}).One(&model.User{}); err == nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "该邮箱已被注册"}
	}
	if err = db.DB(MONGO_DB).C(USER).
		Find(bson.M{"name": u.Name}).One(&model.User{}); err == nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "用户名已存在"}
	}
	must := c.FormValue("2333")
	if must != "on" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "核心价值观呢？这不社会主义！"}
	}

	//Save user
	if err = db.DB(MONGO_DB).C(USER).Insert(u); err != nil {
		return
	}

	cookie := &model.Cookie{
		ID:     u.ID.Hex(),
		Avatar: u.Avatar,
	}
	cookie.WriteCookie(c)

	//u.Password = ""
	return c.Redirect(http.StatusFound, "/")
}

func (h *Handler) Signin(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}
	return c.Render(http.StatusOK, "signin", data)
}

func (h *Handler) Login(c echo.Context) (err error) {
	//清除cookie
	for _, cookie := range c.Cookies() {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}
	//Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB(MONGO_DB).C(USER).
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "用户名或密码错误"}
		}
		return
	}

	cookie := &model.Cookie{
		ID:     u.ID.Hex(),
		Avatar: u.Avatar,
	}
	remember := c.FormValue("remember")
	if remember == "on" {
		cookie.Remember = true
	}
	cookie.WriteCookie(c)

	//u.Password = ""
	return c.Redirect(http.StatusFound, "/")
}

func (h *Handler) Follow(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}
	id := c.Param("id")
	article := c.QueryParam("article")
	pos := c.QueryParam("pos")
	//Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	u := &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(id)).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return err
		}
	}
	if u.IsFollower == nil {
		u.IsFollower = make(map[string]bool)
	}
	if v, ok := u.IsFollower[data.ID]; ok {
		if v {
			u.Follower--
			u.IsFollower[data.ID] = false
		} else {
			u.Follower++
			u.IsFollower[data.ID] = true
		}
	} else {
		u.Follower++
		u.IsFollower[data.ID] = true
	}

	if err = db.DB(MONGO_DB).C(USER).
	UpdateId(u.ID, u); err != nil{
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return err
		}
	}

	if err = db.DB(MONGO_DB).C(USER).
	UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$addToSet": bson.M{"follow": u.ID.Hex()}}); err != nil{
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return err
		}
	}

	return c.Redirect(http.StatusFound, "/article/" + article + "#" + pos)
}

func (h *Handler) Signout(c echo.Context) (err error) {
	for _, cookie := range c.Cookies() {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}
	return c.Redirect(http.StatusFound, "/login")
}

func (h *Handler) UserDetail(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		User *model.User
	}{
		User: &model.User{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}

	id := c.Param("id")
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB(MONGO_DB).C(USER).
	FindId(bson.ObjectIdHex(id)).
	One(data.User); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return err
		}
	}
	data.User.Password = ""
	return c.JSON(http.StatusOK, data.User)
}