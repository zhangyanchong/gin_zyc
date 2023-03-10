package db

import (
	"fmt"
	"gin_zyc/app/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

//初始化数据库
func MysqlInit() {
	appJson, err := tools.ParseConfig("app")
	if err != nil {
		fmt.Println(err)
	}
	//mysql 配置struct
	dbStruct := appJson.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbStruct.Username, dbStruct.Password, dbStruct.Ip, dbStruct.Port, dbStruct.Database, dbStruct.Charset)
	fmt.Println(dsn)
	dbcon, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbcon
}
