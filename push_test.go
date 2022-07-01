package main

import (
	"canal/utils/rabbitmq"
	"fmt"
	"testing"
)

func Test_push(t *testing.T) {
	queueExchange := rabbitmq.QueueExchange{
		"test-queue",
		"test-routingkey",
		"test-exchange",
		"",
		"amqp://guest:guest@127.0.0.1:5672/",
	}

	body := fmt.Sprintf("{\"order_id\":%d}", 320)
	fmt.Println(body)
	_ = rabbitmq.Send(queueExchange, body)
}
