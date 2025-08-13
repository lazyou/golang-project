package config

import "fmt"

type RabbitMQ struct {
	User string `json:"user" yaml:"user"`
	Pass string `json:"pass" yaml:"pass"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func (r *RabbitMQ) Dsn() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", r.User, r.Pass, r.Host, r.Port)
}
