package login

import (
	token2 "easy/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录验证
func Login(c *gin.Context) {
	user := c.PostForm("user")
	password := c.PostForm("password")
	if user == "user" && password == "password" {
		token, _ := token2.CreatToken(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
			"token":   "",
		})
	}
}
