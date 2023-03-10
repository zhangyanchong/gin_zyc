package user

import (
	"fmt"
	"gin_zyc/app/model"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"github.com/gin-gonic/gin"
	"time"
)

type userController struct {
}

var UserController userController

/*用户列表**/
func (a *userController) GetList(c *gin.Context) {
	user := []model.User{}
	db.DB().Find(&user)
	tools.Success(c, "user", user)
}

/*添加用户**/
func (a *userController) Add(c *gin.Context) {
	u := model.User{
		Name:       "wangwu",
		Age:        5,
		Phone:      "14566787551",
		UpdateTime: int(time.Now().Unix()),
		CreateTime: int(time.Now().Unix()),
	}
	rs := db.DB().Create(&u)
	if rs.Error != nil {
		tools.Success(c, "插入失败", nil)
	}
	tools.Success(c, "ok", u)
}

/*修改用户**/
func (a *userController) Update(c *gin.Context) {
	user := model.User{}
	db.DB().Where("id = ?", 3).Take(&user)
	user.Age = 30
	rs := db.DB().Save(&user)
	if rs.Error != nil {
		tools.Success(c, "修改失败", nil)
	}
	tools.Success(c, "ok", nil)
}

/**删除用户*/
func (a *userController) Delete(c *gin.Context) {
	rs := db.DB().Delete(&model.User{}, 7)
	fmt.Println(rs)
	if rs.Error != nil {
		tools.Success(c, "删除失败", nil)
	}
	tools.Success(c, "删除成功", nil)
}
