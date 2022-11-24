package mq

import (
	"fmt"
)

type Producer interface {
	Init()
	Fire()
	Close()
}

type Consumer interface {
	Init()
	Receive()
	Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
