package main

import (
	"go-test/pkg/mq"
)

func main() {
	p := &mq.RabbitMQConsumer{}
	p.Init()
	p.Receive()
	p.Close()
}
