package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"log"
	"net/http"
	"time"
)

func (h *Handler) SignupGet(c echo.Context) (err error) {
	userID, _, _ := userInfoFromToken(c)
	return c.Render(http.StatusOK, "signup", userID)
}

func (h *Handler) Signup(c echo.Context) (err error) {
	//Bind
	u := &model.User{
		ID: bson.NewObjectId(),
	}

	if err = c.Bind(u); err != nil {
		return
	}

	log.Println(u)

	//Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	//Save user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("huanyu0w0").C("users").Insert(u); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Signin(c echo.Context) (err error) {
	userID, _, _ := userInfoFromToken(c)
	return c.Render(http.StatusOK, "signin", userID)
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
	if err = db.DB("huanyu0w0").C("users").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	//JWT
	//Create token
	token := jwt.New(jwt.SigningMethodHS256)

	//Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["name"] = u.Name
	claims["avatar"] = u.Avatar
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()

	//Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return
	}

	u.Password = ""
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Follow(c echo.Context) (err error) {
	userID, _, _ := userInfoFromToken(c)
	id := c.Param("id")

	//Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("huanyu0w0").C("users").
		UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"followers": userID}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}
	return
}

//使用token以及claims的组合来传递当前登录信息更为优雅
func userInfoFromToken(c echo.Context) (id string, name string, avatar string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string), claims["name"].(string), claims["avatar"].(string)
}

