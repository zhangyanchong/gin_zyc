package main

import (
	"fmt"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"gin_zyc/route"
	"github.com/gin-gonic/gin"
)

func main() {

	// 1.创建路由
	r := gin.Default()

	//路由注册
	route.Register(r)

	//数据库连接
	db.MysqlInit()

	appJson, err := tools.ParseConfig("app")
	if err != nil {
		fmt.Println(err)
	}

	r.Run(":" + appJson.AppPort)
}
