package model

type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `form:"username"  binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (Admin) TableName() string {
	return "admin"
}
