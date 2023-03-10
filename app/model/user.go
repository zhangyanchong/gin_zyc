package model

type User struct {
	Id         int
	Name       string
	Age        int
	Phone      string
	UpdateTime int
	CreateTime int
}

func (User) TableName() string {
	return "user"
}

type UserGoods struct {
	UserId    int
	Name      string
	Age       int
	GoodsName string `gorm："column:goods_name"`
}
