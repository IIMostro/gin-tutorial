package configuration

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var channel *amqp.Channel

func getRabbitConnection() *amqp.Channel {

	rabbit := Properties.Rabbit
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbit.Username, rabbit.Password, rabbit.Host, rabbit.Port)
	dial, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("get rabbitmq connection error %f", err)
	}
	session, err := dial.Channel()

	if err != nil {
		log.Fatalf("get rabbitmq connection error %f", err)
	}
	channel = session
	defer session.Close()
	return session
}

func init() {
	channel = getRabbitConnection()
	log.Printf("rabbitmq connection success!")
}
