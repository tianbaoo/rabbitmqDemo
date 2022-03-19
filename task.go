package main

import (
	"rabbitmqDemo/lib"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "no task"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
func main() {
	// 连接RabbitMQ服务器
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "Failed to connect to RabbitMQ")
	// 关闭连接
	defer conn.Close()
	// 新建一个通道
	ch, err := conn.Channel()
	lib.ErrorHanding(err, "Failed to open a channel")
	// 关闭通道
	defer ch.Close()
	// 声明或者创建一个队列用来保存消息
	q, err := ch.QueueDeclare(
		// 队列名称
		"task:queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	lib.ErrorHanding(err, "Failed to declare a queue")
	for i := 0; i < 100000; i++ {
		//body := bodyFrom(a)
		body := fmt.Sprintf("%d", i)
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				// 将消息标记为持久消息
				DeliveryMode: amqp.Persistent,
				Body:         []byte(body),
			})
		log.Printf("sent %s", body)
	}

	//body := bodyFrom(os.Args)
	//err = ch.Publish(
	//	"",
	//	q.Name,
	//	false,
	//	false,
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		// 将消息标记为持久消息
	//		DeliveryMode: amqp.Persistent,
	//		Body:         []byte(body),
	//	})
	lib.ErrorHanding(err, "Failed to publish a message")
	//log.Printf("sent %s", body)
}

