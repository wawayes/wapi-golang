package database

import (
	"fmt"
	"wapi/config"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.GetAddr(),
		Password: config.Conf.Redis.Password, // no password set
		DB:       config.Conf.Redis.DB,       // use default DB
	})

	fmt.Println("[Success] Redis数据库连接成功！！！")
	Client = client
}
