package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostudy/demo/gomod/hello"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		hello.SayHello()
	})
	err := route.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
