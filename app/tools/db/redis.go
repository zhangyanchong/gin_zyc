package db

import (
	"context"
	"fmt"
	"gin_zyc/app/tools"
	"github.com/redis/go-redis/v9"
)

var redisdb *redis.Client

func Redis() *redis.Client {
	return redisdb
}

func RedisInit() {
	appJson, errcfg := tools.ParseConfig("app")
	if errcfg != nil {
		fmt.Println(errcfg)
	}
	//mysql 配置struct
	dbStruct := appJson.Redis
	// 创建
	redisClient := redis.NewClient(&redis.Options{
		Addr:     dbStruct.Ip + ":" + dbStruct.Port,
		Password: dbStruct.Password,
		DB:       dbStruct.Database,
	})
	// 使用超时上下文，验证redis
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
	defer redisClient.Close()
	redisdb = redisClient
}
