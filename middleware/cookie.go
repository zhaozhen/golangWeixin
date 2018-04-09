package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangWeixin/config"
	"golangWeixin/model"
)

// RefreshTokenCookie 刷新过期时间
func RefreshTokenCookie(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	fmt.Println(tokenString)
	if tokenString != "" && err == nil {
		c.SetCookie("token", tokenString, config.ServerConfig.TokenMaxAge, "/", "", true, true)
		if user, err := getUser(c); err == nil {
			model.UserToRedis(user)
		}
	}
	c.Next()
}
