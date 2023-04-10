package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/gin-contrib/sessions"

type userMiddleware struct {
}

var UserMiddleware userMiddleware

func (a *userMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("username") == nil { //没有登录 调整到登录页面
			//fmt.Println(444)
			c.Redirect(http.StatusTemporaryRedirect, "/views/user/login")
		}
		c.Next()
	}
}
