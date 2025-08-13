package initialize

import (
	"github.com/go-redis/redis"
	"xmlt/global"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Pass, // no password set
		DB:       global.Config.Redis.DB,   // use default DB
	})
	ping := client.Ping()
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	global.Redis = client
}
