package admin

import (
	"gin_zyc/app/model"
	"gin_zyc/app/tools"
	"github.com/gin-gonic/gin"
)

type adminController struct {
}

var AdminController adminController

func (a adminController) Login(c *gin.Context) {
	var loginReq model.LoginReq
	if err := c.ShouldBindJSON(loginReq); err != nil {
		tools.Error(c, "参数必填", nil)
		return
	}
	tools.Success(c, "ok", nil)
}
