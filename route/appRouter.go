package route

import "github.com/gin-gonic/gin"
import "gin_zyc/app/controller"

func Register(r *gin.Engine) {
	r.GET("/hello", controller.IndexController.Hello)
	r.GET("/redisSet", controller.RedisController.RedisSet)
	r.GET("/redisGet", controller.RedisController.RedisGet)
	UserRoute.Register(r)
}
