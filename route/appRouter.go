package route

import "github.com/gin-gonic/gin"
import "gin_zyc/app/controller"

func register(r *gin.Engine) {
	r.GET("/hello", controller.IndexController.Hello)
}
