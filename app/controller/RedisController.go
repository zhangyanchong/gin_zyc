package controller

import (
	"context"
	"fmt"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"github.com/gin-gonic/gin"
)

type redisController struct {
}

var RedisController redisController

/*设置redis**/
func (a *redisController) RedisSet(c *gin.Context) {
	_, err := db.Redis().Set(context.Background(), "test", "哈哈哈!", 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	tools.Success(c, "set ok", nil)
}

/*获取redis**/
func (a *redisController) RedisGet(c *gin.Context) {
	rs, err := db.Redis().Get(context.Background(), "test").Result()
	if err != nil {
		fmt.Println(err)
	}
	tools.Success(c, "get ok", rs)
}
