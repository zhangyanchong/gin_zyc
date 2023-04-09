package route

import (
	"gin_zyc/app/controller/html"
	"github.com/gin-gonic/gin"
)

type htmlRoute struct {
}

var HtmlRoute htmlRoute

func (c htmlRoute) Register(r *gin.Engine) {
	r.GET("/views/user/login", html.HtmlController.Login)
	r.GET("/", html.HtmlController.Login)
}
