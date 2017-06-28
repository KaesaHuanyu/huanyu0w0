package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"huanyu0w0/model"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) Topic(c echo.Context) (err error) {
	topic := c.Param("topic")
	if topic != "toys" && topic != "movie" && topic != "tv" && topic != "music_note" &&
		topic != "book" && topic != "others" && topic != "videogame_asset" {
		return c.NoContent(http.StatusNotFound)
	}
	//初始化数据
	data := &struct {
		model.Cookie
		Displays     []*model.Display
		NextPage     int
		PreviousPage int
		Head         bool
		Tail         bool
		Ad           string
		ByTime       bool
		ByLike       bool
	}{
		Displays: []*model.Display{},
		Ad:       "topic/" + topic,
	}
	if err = data.ReadCookie(c); err == nil {
		data.IsLogin = true
	}
	sort := "-" + c.QueryParam("sort")
	if sort == "-" {
		sort = "-like"
	}
	if sort == "-like" {
		data.ByLike = true
	} else if sort == "-time" {
		data.ByTime = true
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	articles := []*model.Article{}
	//Default
	if page == 0 {
		page = 1
	}
	data.NextPage = page + 1
	data.PreviousPage = page - 1
	if data.PreviousPage == 0 {
		data.Head = true
	}

	//查找所有的Display
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB(MONGO_DB).C(ARTICLE).
		Find(bson.M{"topic": topic}).
		Sort(sort).
		Skip((page - 1) * 20).
		Limit(20).
		All(&articles); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	if len(articles) < 20 {
		data.Tail = true
	}

	wg := sync.WaitGroup{}
	for i, v := range articles {
		wg.Add(1)
		data.Displays = append(data.Displays, &model.Display{})
		data.Displays[i].Article = v
		data.Displays[i].ID = v.ID.Hex()
		go func(i int) {
			defer wg.Done()
			data.Displays[i].ShowTime = data.Displays[i].Article.GetShowTime()
			data.Displays[i].ShowTopic = data.Displays[i].Article.GetShowTopic()
			data.Displays[i].CommentsNum = len(data.Displays[i].Article.Comments)
			//data.Displays[i].Editor不能为 nil
			data.Displays[i].Editor = &model.User{}
			db := h.DB.Clone()
			defer db.Close()
			err = db.DB(MONGO_DB).C(USER).
				FindId(bson.ObjectIdHex(data.Displays[i].Article.Editor)).
				One(data.Displays[i].Editor)
			if err != nil {
				fmt.Println("<(￣︶￣)↗[GO!]", i, ":", err)
			}
			data.Displays[i].Editor.Password = ""
		}(i)
	}
	wg.Wait()

	return c.Render(http.StatusOK, "home", data)
}
