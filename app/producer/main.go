package main

import (
	"go-test/mq"
)

func main() {
	p := &mq.RabbitMQProducer{}
	p.Init()
	msgs := []string{"test1", "test2", "test3", "test4", "test5"}
	for _, msg := range msgs {
		p.Fire(msg)
	}
	p.Close()
}