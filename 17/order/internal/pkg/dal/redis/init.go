package redis

import goredis "github.com/go-redis/redis"

var client *goredis.Client

func init() {
	// 创建 Redis 客户端
	client = goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
