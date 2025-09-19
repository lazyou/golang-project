package config

import (
	"go-gin/internal/component/db"
	"go-gin/internal/component/logx"
	"go-gin/internal/component/redisx"
	filex "go-gin/internal/file"
	"sync"
)

type Config struct {
	App     App           `yaml:"app"`
	Redis   redisx.Config `yaml:"redis"`
	DB      db.Config     `yaml:"db"`
	Log     logx.Config   `yaml:"log"`
	Svc     SvcConfig     `yaml:"svc"`
	Monitor MonitorConfig `yaml:"monitor"`
}

var instance *Config
var once sync.Once

func Init(filename string) {
	// 【单例|初始化配置】sync.Once 包实现单例、或者配置文件初始化，避免二次覆盖
	once.Do(func() {
		err := filex.MustLoad(filename, &instance)
		if err != nil {
			panic(err)
		}
	})
}
