package main

import (
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/controller"
	"gostudy/answer/single-arch/model"
)

var engine *gin.Engine

func main() {
	engine = gin.Default()
	HandleStatic()
	//用户注册
	engine.POST("/user/register", (&controller.UserController{}).Register)
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
}
