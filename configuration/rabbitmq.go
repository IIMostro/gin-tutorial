package configuration

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func GetRabbitConnection() *amqp.Channel {

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
	return session
}
