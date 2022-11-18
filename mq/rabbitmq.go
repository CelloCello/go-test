package mq

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	ch    *amqp.Channel
	queue *amqp.Queue
	conn  *amqp.Connection
}

type RabbitMQConsumer struct {
	ch    *amqp.Channel
	queue *amqp.Queue
	conn  *amqp.Connection
}

func (p *RabbitMQProducer) Init() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	p.conn = conn

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()
	p.ch = ch

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
	p.queue = &q

}

func (p *RabbitMQProducer) Fire(msg string) {
	body := msg
	err := p.ch.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Printf(" [x] Sent %s\n", body)
}

func (p *RabbitMQProducer) Close() {
	p.ch.Close()
	p.conn.Close()
}

func (c *RabbitMQConsumer) Init() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	c.conn = conn

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()
	c.ch = ch

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
	c.queue = &q
}

func (c *RabbitMQConsumer) Receive() {
	msgs, err := c.ch.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (c *RabbitMQConsumer) Close() {
	c.ch.Close()
	c.conn.Close()
}