package middleware

import (
	token2 "easy/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token2.IsValid(c) {
			// 认证成功，可以将用户信息保存在上下文中
			c.Next()
		} else {
			// 认证失败，响应未认证
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}
	}
}
