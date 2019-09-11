package main

import (
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/controller"
	"gostudy/answer/single-arch/filter"
	"gostudy/answer/single-arch/id_gen"
	"gostudy/answer/single-arch/middleware/account"
	"gostudy/answer/single-arch/model"
)

var engine *gin.Engine

func main() {
	engine = gin.Default()
	HandleStatic()
	//用户模块
	user := engine.Group("/api/user")
	{
		//用户注册
		user.POST("/register", (&controller.UserController{}).Register)
		//用户登录
		user.POST("/login", (&controller.UserController{}).Login)
	}
	engine.GET("/api/category/list", (&controller.CategoryController{}).GetCategoryList)
	//中间件使用
	engine.POST("/api/ask/submit", account.AuthMiddleware, (&controller.QuestionController{}).AddQuestion)

	err := engine.Run(":9090")
	if err != nil {
		panic(err)
	}
}

func HandleStatic() {
	engine.StaticFile("/", "./static/index.html")
	engine.Static("/css", "./static/css")
	engine.Static("/img", "./static/img")
	engine.Static("/js", "./static/js")
	engine.Static("/fonts", "./static/fonts")
	engine.StaticFile("/favicon.ico", "./static/favicon.ico")
}

func init() {
	model.Init()
	id_gen.Init(0)
	account.Init()
	err := filter.Init("./filter/words.txt")
	if err != nil {
		panic(err)
	}
}
