package token

import (
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
	paseto2 "github.com/o1egl/paseto/v2"
	"time"
)

var symmetricKey = []byte("YELLOW SUBMARINE, BLACK WIZARDRY") // Must be 32 bytes

func IsValid(c *gin.Context) bool {
	token := getToken(c)
	JToken, _ := parseToken(token)
	return checkToken(JToken)
}

// CreatToken 生成token
func CreatToken(user string) (string, error) {
	jsonToken := paseto.JSONToken{
		Audience:   user,                             //接受token的用户
		Issuer:     "paseto",                         //颁发者
		Jti:        "nobody guess it",                //Token ID(尽量保证id唯一)
		Subject:    "subject",                        //主题
		IssuedAt:   time.Now(),                       //颁发时间
		Expiration: time.Now().Add(20 * time.Second), //过期时间
		NotBefore:  time.Now(),                       //生效时间
	}

	// 添加自定义信息
	//jsonToken.Set("data", "this is a signed message")

	// Encrypt
	return paseto2.Encrypt(symmetricKey, jsonToken, nil)
}

// 验证token
func checkToken(jsonToken paseto.JSONToken) bool {
	if jsonToken.Audience == "user" && jsonToken.Expiration.After(time.Now()) && jsonToken.NotBefore.Before(time.Now()) {
		return true
	} else {
		return false
	}
}

// 解析token
func parseToken(token string) (paseto.JSONToken, error) {
	// Decrypt
	var JToken paseto.JSONToken
	return JToken, paseto2.Decrypt(token, symmetricKey, &JToken, nil)
}

// 获取token
func getToken(c *gin.Context) string {
	return c.Request.Header.Get("Token")
	//return c.Request.Header["Token"][0]
}
