package config

import (
	"go-gin/internal/component/db"
	"go-gin/internal/component/logx"
	"go-gin/internal/component/redisx"
)

// 配置项的统一获取和分开获取

func GetAppConf() App {
	return instance.App
}

func GetRedisConf() redisx.Config {
	return instance.Redis
}

func GetLogConf() logx.Config {
	return instance.Log
}

func GetDbConf() db.Config {
	return instance.DB
}

func GetMonitorConf() MonitorConfig {
	return instance.Monitor
}
