package route

import (
	"gin_zyc/app/controller/html"
	"gin_zyc/app/middleware"
	"github.com/gin-gonic/gin"
)

type htmlRoute struct {
}

var HtmlRoute htmlRoute

func (c htmlRoute) Register(r *gin.Engine) {
	r.GET("/views/user/login", html.HtmlController.Login)
	r.GET("/", html.HtmlController.Login)

	//登录中间件使用
	loginAuth := r.Group("/", middleware.UserMiddleware.AuthRequired())
	loginAuth.GET("/views/blog/list", html.HtmlController.BlogList)
	loginAuth.GET("/views/blog/add", html.HtmlController.BlogAdd)
	loginAuth.GET("/views/blog/detail", html.HtmlController.BlogDetail)

	loginAuth.POST("/api/blog/blogAddApi", html.HtmlController.BlogAddApi)
	loginAuth.POST("/api/blog/blogDelApi", html.HtmlController.BlogDelApi)

}
