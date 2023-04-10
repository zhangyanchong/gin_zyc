package main

import (
	"fmt"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"gin_zyc/route"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// 1.创建路由
	r := gin.Default()

	//设置session
	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("zyc"))
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	//加载静态资源
	r.Static("/assets", "./assets")
	//加载html 页面
	r.LoadHTMLGlob("views/*/*")

	//路由注册
	route.Register(r)

	//数据库连接
	db.MysqlInit()

	//redis连接
	db.RedisInit()

	appJson, err := tools.ParseConfig("app")
	if err != nil {
		fmt.Println(err)
	}

	r.Run(":" + appJson.AppPort)
}
