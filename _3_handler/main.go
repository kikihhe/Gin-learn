package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	TOKEN = "123"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != TOKEN {
			c.Abort()
		}
		c.Set("id", "xiaohe")
	}
}

func GetUser(c *gin.Context) {
	fmt.Println("访问 /getUser")
	if id, ok := c.Get("id"); ok {
		c.String(200, "成功返回数据, id:", id)
	} else {
		c.String(500, "错误, id为空")
	}
}

// 将api访问记录插入数据库
func addApiRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Get("id")
		fmt.Printf("%s 开始访问系统!\n", id)
		ctx.Abort()
		ctx.Next()
		fmt.Printf("结束访问系统")
	}
}
func main() {
	r := gin.Default()

	r.Use(CheckToken())
	r.Use(addApiRecord())
	{
		r.GET("/getUser", GetUser)
	}

	r.Run(":8080")
}
