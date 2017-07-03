package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"huanyu0w0/model"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) Home(c echo.Context) (err error) {
	//初始化数据
	data := &struct {
		model.Cookie
		Displays     []*model.Display
		NextPage     int
		PreviousPage int
		Head         bool
		Tail         bool
		ByLike       bool
		ByTime       bool
		Ad           string
	}{
		Displays: []*model.Display{},
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
	////Default
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
		Find(nil).
		Sort(sort).
		Skip((page - 1) * 12).
		Limit(12).
		All(&articles); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	if len(articles) < 12 {
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
			err := db.DB(MONGO_DB).C(USER).
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

func (h *Handler) CurriculumVitae(c echo.Context) error {
	data := struct {
		One   template.HTML
		Two   template.HTML
		Three template.HTML
		Four  template.HTML
		Five  template.HTML
		Six   template.HTML
	}{}

	data.One = template.HTML(blackfriday.MarkdownCommon([]byte(`##### 姓名：黎寰宇`)))

	data.Two = template.HTML(blackfriday.MarkdownCommon([]byte(`#### 四川华迪信息技术有限公司-研发实习生（2016-07-14 ~ 2016-08-15）
#### 项目经历
与同学组成小组搭建了一个小型的电商网站，包括用户登录注册，商品展示管理，用户下单至购买的整套流程。

我在其中负责商品部分前端页面的编写（包括商品分类展示的多个页面，商品详情页面等），以及商品部分的后端代码编写（商品的CRUD，加入购物车，商品的排列方式等，涉及与mysql数据库交互）。

---

#### DaoCloud-CTO Office-研发实习生（2017-02-27 ~ 至今）

#### 项目经历
##### 2017-03 ~ 2017-04 上汽saiccloud项目第一期
我所在的部门在本项目中负责服务目录的容器化解决方案，本期计划上线的共有四个服务：mysql高可用服务，mongoDB高可用服务，RabbitMQ高可用服务，Redis高可用服务。

我的主要工作为

- 测试这些服务在部署之后能否按预期工作，并在测试发现问题后解决问题（重写Dockerfile、镜像启动脚本等）。
- 编写代码测试或者解决来自上级的需求，编写服务高可用测试程序并做成了公开镜像 daocloud.io/daocloud/servicetest
- 服务交付文档、服务测试文档以及其他需要交付的文档的编写

##### 2017-04 ~ 2017-05 上汽saiccloud项目第二期
在第二期项目中，客户追加了对zookeeper、kafka、tomcat的需求。

我的工作与之前基本一致（但是这一期中zookeeper和kafka基本交于我负责）

具体成果：

- 发现并单独解决了zk不能正常建立集群，kafka的broker不能访问外网、
kafka不能正常连接到zookeeper等严重问题。
- 独自完成keepalived的自动化配置虚拟ip的解决方案
- 各类文档以及测试报告

---
#### 个人项目
##### 2017-04 ~ 至今： [huanyu0w0](https://github.com/KaesaHuanyu/huanyu0w0)网站的搭建
本项目使用的工具为:

- 后端：由golang编写，使用[echo](https://echo.labstack.com/guide)框架
- 前端：使用谷歌公司的[MDL](https://getmdl.io/index.html)框架
- 数据库：使用[mongoDB](https://hub.docker.com/_/mongo/)
- 部署方式：在阿里云主机上使用[docker](https://www.docker.com/)部署。
- 域名 www.huanyu0w0.cn(已备案)
- 图床：七牛云存储
- MD编辑器：[Editor.md](https://github.com/pandao/editor.md)

网站分为文章、评论、用户三大模块，主要用例有：用户注册登录，文章编写查看，评论文章，回复评论，点赞功能等。`)))

	data.Three = template.HTML(blackfriday.MarkdownCommon([]byte(`
- 熟悉golang
- 熟悉docker的日常使用
- 熟悉linux的日常使用
- 熟悉git的日常使用
- 了解常用DB、MQ及其高可用架构
- 了解前端，熟悉后端
- 基础知识良好，熟悉常用算法和数据结构
- 能够较为熟练查阅英文技术文档，学习能力强`)))

	data.Four = template.HTML(blackfriday.MarkdownCommon([]byte(`
 - 个人网站：[www.huanyu0w0.cn](http://www.huanyu.cn)

 - GitHub：[https://github.com/KaesaHuanyu](https://github.com/KaesaHuanyu)
 - 爱好Label：二次元、音乐、日剧、lol、科幻小说（间客、地球纪元、三体之类）
 - 梦想：在一个安静的街角开一间咖啡厅
 - 座右铭：己所不欲勿施于人`)))

	data.Five = template.HTML(blackfriday.MarkdownCommon([]byte(`
- 学校：电子科技大学
- 专业：软件工程
- 学历：本科
- 年级：2014级
- GPA：2.88/4.0
- 排名：专业37名（前30%）
- 曾获奖学金：人民奖学金`)))

	data.Six = template.HTML(blackfriday.MarkdownCommon([]byte(`
- 手机号码：13990591066
- 邮箱：kaesalai@gmail.com
- QQ：875386471`)))

	return c.Render(http.StatusOK, "curriculumVitae", data)
}
