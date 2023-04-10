package admin

import (
	"gin_zyc/app/model"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type adminController struct {
}

var AdminController adminController

func (a adminController) Login(c *gin.Context) {
	var loginReq model.LoginReq
	if err := c.ShouldBind(&loginReq); err != nil {
		tools.Error(c, "参数必填", nil)
		return
	}
	admin := model.Admin{}
	db.DB().Where("username=? and password=?", loginReq.Username, loginReq.Password).Find(&admin)

	//成功登录设置session
	if admin.Id > 0 {
		session := sessions.Default(c)
		session.Set("username", admin.Username)
		session.Set("user_id", admin.Id)
		session.Save()
	}

	tools.Success(c, "ok", admin)
}
