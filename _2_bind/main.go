package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

func InsertUser(context *gin.Context) {
	var user UserInfo
	err := context.ShouldBind(&user)

	if err != nil {
		context.String(500, "参数错误")
	}

	fmt.Println("插入数据: ", user)
	context.String(200, "插入成功")
}

type Head struct {
	Token string
	Aaa   string
}

func BindHeaderTest(context *gin.Context) {
	var header Head
	err := context.ShouldBindHeader(&header)
	if err != nil {
		context.String(500, "参数错误")
	}
	fmt.Println("获取请求头部数据: ", header)
	context.String(200, "获取成功")
}

type URI struct {
	Id   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required,max=6" msg:"name最短为3最长为6"`
}

func BindUrlTest(context *gin.Context) {
	var uri URI
	err := context.ShouldBindUri(&uri)
	if err != nil {
		context.String(500, "参数错误")
		context.Abort()
	}
	fmt.Println("获取uri数据: ", uri)
	context.String(200, "获取成功")
}

func main() {
	r := gin.Default()

	r.POST("/user/insert", InsertUser)
	r.POST("/user/getHeader", BindHeaderTest)
	r.POST("/user/:id/:name", BindUrlTest)

	r.Run(":8080")
}
