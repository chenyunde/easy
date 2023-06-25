package main

import (
	"easy/login"
	"easy/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 权限验证中间件
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())

	// 不需要权限验证
	noAuth := r.Group("/noauth")
	noAuth.POST("/login", login.Login)

	r.Run(":8080")
}
