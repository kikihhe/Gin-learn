package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context) {

}
func GetUserByName(c *gin.Context) {

}
func GetUserByAge(c *gin.Context) {

}
func GetDeviceById(c *gin.Context) {

}

func main() {
	r := gin.Default()
	// 全局中间件跨域
	r.Use(Cors())
	// 用户
	userGroup := r.Group("/user")
	{
		userGroup.GET("/getUserById", GetUserById)
		userGroup.GET("/getUserByName", GetUserByName)
		userGroup.GET("/getUserByAge", GetUserByAge)
	}

	// 设备
	deviceGroup := r.Group("/device")
	{
		deviceGroup.GET("getDeviceById", GetDeviceById)
	}
}

// Cors 跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有的域名进行跨域请求
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许跨域的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		// 允许跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		// 允许跨域请求的响应头
		c.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers, Content-Type")
		// 允许跨域请求携带cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		method := c.Request.Method
		// OPTIONS方法用于获取目的资源所支持的通信选项
		if method == "OPTIONS" {
			// 如果是OPTIONS请求，直接返回空响应
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
