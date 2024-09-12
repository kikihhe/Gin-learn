package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(context *gin.Context) {
	context.String(200, "hello world!")
}

func GetQueryString(context *gin.Context) {
	name := context.Query("name")
	age := context.DefaultQuery("age", "0")
	userMap := context.QueryMap("user")
	fmt.Println(userMap)
	fmt.Println("name:", name)
	fmt.Printf("age的类型: %T\n", age)
}

func GetPostForm(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	c.String(http.StatusOK, fmt.Sprintf("username: %s, pwd: %s, type: %s", username, pwd, types))
}
func GetPath(context *gin.Context) {
	name := context.Param("name")
	age := context.Param("age")
	context.String(200, "name:%s, age:%s", name, age)
}
func GetPostJson(ctx *gin.Context) {
	// 获取单个值
	name := ctx.PostForm("name")
	//带默认值
	gender := ctx.DefaultPostForm("gender", "男")
	//数组
	habits := ctx.PostFormArray("habits")
	//map
	works := ctx.PostFormMap("works")
	fmt.Printf("%s,%s,%s,%s\n", name, gender, habits, works)

}
func GetHeader(c *gin.Context) {
	token := c.GetHeader("token")
	fmt.Println(token)
}

func main() {
	// 创建路由
	r := gin.Default()

	r.GET("/helloWorld", HelloWorld)
	r.GET("/getQueryString", GetQueryString)
	r.GET("/getPath/:name/:age", GetPath)
	r.GET("/getHeader", GetHeader)

	r.Run(":8080")
}
