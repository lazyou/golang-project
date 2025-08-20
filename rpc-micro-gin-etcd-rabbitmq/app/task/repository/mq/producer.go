package mq

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/CocaineCong/micro-todoList/consts"
)

// SendMessage2MQ 发送消息到mq
func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}

	q, _ := ch.QueueDeclare(consts.RabbitMqTaskQueue, true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return
	}

	fmt.Println("生产者-【RabbitMQ 消息入队列成功】...")
	return
}
