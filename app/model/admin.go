package model

type Admin struct {
	Id       int
	Username string
	Password string
}

type LoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (Admin) TableName() string {
	return "admin"
}
