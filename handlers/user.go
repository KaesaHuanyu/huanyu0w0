package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/labstack/echo-contrib/session"
)

func (h *Handler) SignupGet(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Path string
	}{
		Path: c.QueryParam("path"),
	}
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

	verifyUser := c.FormValue("verify")
	if verifyUser == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "请输入验证码"}
	}
	sess, _ := session.Get("session", c)
	if verifyUser != sess.Values["verify"] {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "验证码错误"}
	}

	//Bind
	u := &model.User{
		ID:     bson.NewObjectId(),
		Time:   time.Now(),
		Avatar: "http://images.huanyu0w0.cn/avatar/squirrelAvatar.jpg",
	}

	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}
	//Validate
	if u.Email == "" || u.Password == "" || u.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "不能为空"}
	}

	if len(u.Name) > 24 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "用户名不能超过24个字节"}
	}

	if len(u.Name) > 30 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "邮箱有这么长吗..."}
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
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	cookie := &model.Cookie{
		ID:     u.ID.Hex(),
		Avatar: u.Avatar,
	}
	cookie.WriteCookie(c)

	//记录日志
	log := &model.Log{
		ID:        bson.NewObjectId(),
		Time:      time.Now(),
		User:      u.ID.Hex(),
		Operation: "注册",
		Signup:    true,
	}

	if err = db.DB(MONGO_DB).C(LOG).
		Insert(log); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	//u.Password = ""
	path := c.QueryParam("path")
	return c.Redirect(http.StatusFound, "/"+path)
}

func (h *Handler) Signin(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		Path string
	}{
		Path: c.QueryParam("path"),
	}
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
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB(MONGO_DB).C(USER).
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "用户名或密码错误"}
		}
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}

	cookie := &model.Cookie{
		ID:       u.ID.Hex(),
		Avatar:   u.Avatar,
		Email:    u.Email,
		UnixTime: strconv.Itoa(int(u.Time.Unix())),
		Name:     u.Name,
	}
	remember := c.FormValue("remember")
	if remember == "on" {
		cookie.Remember = true
	}
	cookie.WriteCookie(c)

	//u.Password = ""
	path := c.QueryParam("path")
	return c.Redirect(http.StatusFound, "/"+path)
}

func (h *Handler) Follow(c echo.Context) (err error) {
	article := c.QueryParam("article")
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path="+article)
	}
	id := c.Param("id")
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
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
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
		UpdateId(u.ID, u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	if err = db.DB(MONGO_DB).C(USER).
		UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$addToSet": bson.M{"follow": u.ID.Hex()}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	return c.Redirect(http.StatusFound, "/article/"+article+"#"+pos)
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
		UserDisplay *model.UserDisplay
	}{
		UserDisplay: &model.UserDisplay{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}

	id := c.Param("id")
	db := h.DB.Clone()
	defer db.Close()
	data.UserDisplay.User = &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(id)).
		One(data.UserDisplay.User); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	data.UserDisplay.ID = id
	data.UserDisplay.CreateTime = data.UserDisplay.User.GetCreateTime()
	data.UserDisplay.User.BigAvatar = strings.TrimSuffix(data.UserDisplay.User.Avatar, "-avatarStyle")
	data.UserDisplay.User.Password = ""

	//Articles
	wg := &sync.WaitGroup{}
	for i, v := range data.UserDisplay.User.Articles {
		wg.Add(1)
		data.UserDisplay.Articles = append(data.UserDisplay.Articles, &model.Article{})
		go func(i int, v string) {
			defer wg.Done()
			db := h.DB.Clone()
			defer db.Close()
			if err = db.DB(MONGO_DB).C(ARTICLE).
				FindId(bson.ObjectIdHex(v)).
				One(data.UserDisplay.Articles[i]); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, ", GetArticle:", err)
			}
			data.UserDisplay.Articles[i].ShowID = data.UserDisplay.Articles[i].ID.Hex()
		}(i, v)
	}
	wg.Wait()

	//Follow
	wg2 := &sync.WaitGroup{}
	for i, v := range data.UserDisplay.User.Follows {
		wg2.Add(1)
		data.UserDisplay.Follow = append(data.UserDisplay.Follow, &model.User{})
		go func(i int, v string) {
			defer wg2.Done()
			db := h.DB.Clone()
			defer db.Close()
			if err = db.DB(MONGO_DB).C(USER).
				FindId(bson.ObjectIdHex(v)).
				One(data.UserDisplay.Follow[i]); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, ", GetFollow:", err)
			}
			data.UserDisplay.Follow[i].ShowID = data.UserDisplay.Follow[i].ID.Hex()
			data.UserDisplay.Follow[i].BigAvatar = strings.TrimSuffix(data.UserDisplay.Follow[i].Avatar,
				"-avatarStyle")
			data.UserDisplay.Follow[i].Password = ""
		}(i, v)
	}
	wg2.Wait()

	return c.Render(http.StatusOK, "user", data)
}

func (h *Handler) Dashboard(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
		UserDisplay *model.UserDisplay
	}{
		UserDisplay: &model.UserDisplay{},
	}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path=user/dashboard")
	}

	//取得mongo连接
	db := h.DB.Clone()
	defer db.Close()

	//得到User
	data.UserDisplay.User = &model.User{}
	if err = db.DB(MONGO_DB).C(USER).
		FindId(bson.ObjectIdHex(data.ID)).
		One(data.UserDisplay.User); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		} else {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
		}
	}

	data.UserDisplay.User.BigAvatar = strings.TrimSuffix(data.UserDisplay.User.Avatar, "-avatarStyle")
	data.UserDisplay.User.Password = ""

	//得到Articles
	wg := &sync.WaitGroup{}
	for i, v := range data.UserDisplay.User.Articles {
		wg.Add(1)
		data.UserDisplay.Articles = append(data.UserDisplay.Articles, &model.Article{})
		go func(i int, v string) {
			defer wg.Done()
			db := h.DB.Clone()
			defer db.Close()
			if err = db.DB(MONGO_DB).C(ARTICLE).
				FindId(bson.ObjectIdHex(v)).
				One(data.UserDisplay.Articles[i]); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, ", GetArticle:", err)
			}
			data.UserDisplay.Articles[i].ShowID = data.UserDisplay.Articles[i].ID.Hex()
		}(i, v)
	}
	wg.Wait()

	//得到Comments
	wg2 := &sync.WaitGroup{}
	for i, v := range data.UserDisplay.User.Comments {
		wg2.Add(1)
		data.UserDisplay.Comments = append(data.UserDisplay.Comments, &model.Comment{})
		go func(i int, v string) {
			defer wg2.Done()
			db := h.DB.Clone()
			defer db.Close()
			if err = db.DB(MONGO_DB).C(COMMENT).
				FindId(bson.ObjectIdHex(v)).
				One(data.UserDisplay.Comments[i]); err != nil {
				log.Println("<(￣︶￣)↗[GO!]", i, ", GetComment:", err)
			}
			data.UserDisplay.Comments[i].ShowID = data.UserDisplay.Comments[i].ID.Hex()
		}(i, v)
	}
	wg2.Wait()

	return c.Render(http.StatusOK, "dashboard", data)
}

func (h *Handler) RemoveUser(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		log.Println("Not Login")
		return c.NoContent(http.StatusNotFound)
	}
	return c.Redirect(http.StatusFound, "/user/dashboard")
}

func (h *Handler) UpdateUser(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		log.Println("Not Login")
		return c.NoContent(http.StatusNotFound)
	}

	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s", err)}
	}
	u.Change = time.Now()

	db := h.DB.Clone()
	defer db.Close()

	if u.Name != "" {
		if err = db.DB(MONGO_DB).C(USER).
			Find(bson.M{"name": u.Name}).One(&model.User{}); err == nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "用户名已存在"}
		}
		if err = db.DB(MONGO_DB).C(USER).
			UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$set": bson.M{"name": u.Name}}); err != nil {
			if err == mgo.ErrNotFound {
				return echo.ErrNotFound
			} else {
				return &echo.HTTPError{Code: http.StatusBadRequest, Message: "更新信息失败"}
			}
		}
		if len(u.Name) > 24 {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "用户名不能超过24个字节"}
		}
	}

	if u.Password != "" {
		if err = db.DB(MONGO_DB).C(USER).
			UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$set": bson.M{"password": u.Password}}); err != nil {
			if err == mgo.ErrNotFound {
				return echo.ErrNotFound
			} else {
				return &echo.HTTPError{Code: http.StatusBadRequest, Message: "更新信息失败"}
			}
		}
	}

	if u.Info != "" {
		if err = db.DB(MONGO_DB).C(USER).
			UpdateId(bson.ObjectIdHex(data.ID), bson.M{"$set": bson.M{"info": u.Info}}); err != nil {
			if err == mgo.ErrNotFound {
				return echo.ErrNotFound
			} else {
				return &echo.HTTPError{Code: http.StatusBadRequest, Message: "更新信息失败"}
			}
		}
	}

	return c.Redirect(http.StatusFound, "/user/dashboard")
}
