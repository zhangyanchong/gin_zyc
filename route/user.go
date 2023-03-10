package route

import (
	"gin_zyc/app/controller/user"
	"github.com/gin-gonic/gin"
)

type userRoute struct {
}

var UserRoute userRoute

func (c userRoute) Register(r *gin.Engine) {
	r.GET("/user/getlist", user.UserController.GetList)
	r.GET("/user/add", user.UserController.Add)
	r.GET("/user/update", user.UserController.Update)
	r.GET("/user/delete", user.UserController.Delete)

}
