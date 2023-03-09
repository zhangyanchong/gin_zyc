package main

import (
	"fmt"
	"gin_zyc/app/tools"
	"github.com/gin-gonic/gin"
)

func main() {

	// 1.创建路由
	r := gin.Default()

	appJson, err := tools.ParseConfig("app")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(appJson["app_port"])
	// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})

	r.Run(":" + appJson["app_port"].(string))
}
