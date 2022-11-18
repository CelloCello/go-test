package main

import (
	"go-test/mq"
)

func main() {
	p := &mq.RabbitMQConsumer{}
	p.Init()
	p.Receive()
	p.Close()
}