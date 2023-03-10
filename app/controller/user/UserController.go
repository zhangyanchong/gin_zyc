package user

import (
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
	//修改返回来的结果，换成新的切片操作
	type xinUser struct {
		User model.User
		Sex  int
	}
	var showXinUser []xinUser
	for _, v := range user {
		showXinUser = append(showXinUser, xinUser{User: v, Sex: 20})
	}
	tools.Success(c, "user", showXinUser)
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
	rs := db.DB().Delete(&model.User{}, 6)
	if rs.Error != nil {
		tools.Success(c, "删除失败", nil)
	}
	//判断是否删除成功了
	if rs.RowsAffected > 0 {
		tools.Success(c, "删除成功", nil)
	} else {
		tools.Success(c, "删除失败", nil)
	}
}

/*用户的商品**/
func (a *userController) UserGoods(c *gin.Context) {
	//sqlStr := "select  goods.user_id,user.age,user.name,goods.goods_name from user " +
	//	"left join goods   on user.id=goods.user_id " +
	//	"where  user.age>10"
	sqlStr := "select  b.user_id,a.age,a.name,b.goods_name from user as  a " +
		"left join goods as b  on a.id=b.user_id " +
		"where  a.age>10"
	f := []model.UserGoods{}
	db.DB().Raw(sqlStr).Scan(&f)
	tools.Success(c, "usergoods", f)
}
