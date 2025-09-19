package redisx

import (
	"context"
	"go-gin/internal/component/logx"

	"github.com/redis/go-redis/v9"
)

var (
	instance *redis.Client
	conf     Config
)

type Config struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"` // no password set
	DB       int    `yaml:"db"`       // use default DB
}

func InitConfig(c Config) {
	conf = c
}

func Init() {
	options := &redis.Options{
		Addr:     conf.Addr,
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	}
	rdb := redis.NewClient(options)
	// 向 Redis 客户端注册“钩子”（Hook），在命令执行前后、出错时等时机执行自定义代码。
	// 是一种【AOP（面向切面编程）】 思想的体现：不侵入核心逻辑，但能监控或增强行为。
	// 【重要】ProcessHook：拦截每个 Redis 命令: 记录命令执行时间、记录执行的命令和参数、统计错误率、上报监控系统
	rdb.AddHook(&LogHook{})
	rdb.AddHook(&ErrHook{})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		logx.WithContext(context.Background()).Error("redis", err)
	}
	instance = rdb
}

func Client() *redis.Client {
	return instance
}
