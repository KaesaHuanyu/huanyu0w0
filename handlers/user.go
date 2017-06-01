package handlers

import (
	"github.com/dgrijalva/jwt-go"
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
	return c.JSON(http.StatusCreated, u)
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

	//JWT
	//Create token
	token := jwt.New(jwt.SigningMethodHS256)

	//Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()

	//Generate encoded token and send it as response
	//u.Token, err = token.SignedString([]byte(KEY))
	//if err != nil {
	//	return
	//}

	u.Password = ""
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Follow(c echo.Context) (err error) {
	userID := userInfoFromToken(c)
	id := c.Param("id")

	//Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"followers": userID}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}
	return
}

func (h *Handler) Signout(c echo.Context) (err error) {
	for _, cookie := range c.Cookies() {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}
	return c.Redirect(http.StatusFound, "/login")
}

//使用token以及claims的组合来传递当前登录信息更为优雅
func userInfoFromToken(c echo.Context) (id string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func (h *Handler) UserDetail(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, nil)
}