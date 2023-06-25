package login

import (
	token2 "easy/token"
	"easy/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录验证
func Login(c *gin.Context) {
	username := c.PostForm("userInfo")
	password := c.PostForm("password")
	var u user.UInfo
	u.Init()
	u.UserName = username
	u.PassWord = password
	if u.UInfoCheck() {
		token, _ := token2.CreatToken(username)
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
