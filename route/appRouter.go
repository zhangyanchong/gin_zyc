package route

import "github.com/gin-gonic/gin"
import "gin_zyc/app/controller"
import "gin_zyc/app/controller/admin"

func Register(r *gin.Engine) {
	r.GET("/hello", controller.IndexController.Hello)
	r.GET("/redisSet", controller.RedisController.RedisSet)
	r.GET("/redisGet", controller.RedisController.RedisGet)

	r.POST("/admin/admin/login", admin.AdminController.Login)

	UserRoute.Register(r)
	HtmlRoute.Register(r)
}
