package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	//一、sample
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//二、resetful api
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	//三、路由参数
	//1 :username必须
	router.GET("/routeParam1/:username", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ruote":    "/routeParam/:username",
			"username": c.Param("username"),
		})
	})
	// 2、:username必须，*age可有可无  匹配/routeParam2/snailzed/ 要带上/
	router.GET("/routeParam2/:username/*age", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ruote":    "/routeParam2/:username/*age",
			"username": c.Param("username"),
			"age":      c.Param("age"),
		})
	})

	//四、querystring参数
	router.GET("/queryString", func(c *gin.Context) {
		//c.QueryMap("ids")  /queryString?ids[1]=1&ids[2]=2
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(200, "Welcome! %s %s", firstname, lastname)
	})

	//五、form表单参数解析
	router.POST("/post_form", func(c *gin.Context) {
		//c.PostFormMap("hobbies")  hobbies[]="a"&
		username := c.PostForm("username")
		age := c.DefaultPostForm("age", "24")
		hobbies := c.PostFormMap("hobbies")
		h := c.PostFormArray("hobbies")
		c.String(http.StatusOK, "%s's age is %s,hobbies:%#v,%#v", username, age, hobbies, h)
	})

	//六、文件上传
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source 单文件上传
		//多文件上传：form, err := c.MultipartForm() ; 遍历form.File["files"]
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	})

	//七、分组路由
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		//路由是：/v2/login
		v2.POST("/login", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
	}

	err := router.Run(":8085") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}

	//八、渲染输出
	// c.JSON(code,data interface{}) data可以是任意类型的，会自动序列化
	// c.XML(code int, data interface{})
	// 	router.Static() : 设置静态文件路径

}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "GET",
	})
}

func posting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "POST",
	})
}

func putting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "PUT",
	})
}
func deleting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "DELETE",
	})
}
func patching(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "PATCH",
	})
}
func head(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "HEAD",
	})
}
func options(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "OPTIONS",
	})
}
