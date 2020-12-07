package configuration

import (
	"fmt"
	"github.com/streadway/amqp"
)

func GetRabbitConnection() *amqp.Channel {

	rabbit := Properties.Rabbit
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbit.Username, rabbit.Password, rabbit.Host, rabbit.Port)
	dial, err := amqp.Dial(url)
	if err != nil {
		err := fmt.Errorf("get rabbitmq connection error %w", err)
		panic("get rabbitmq error, cause: " + err.Error())
	}
	session, err := dial.Channel()

	if err != nil {
		err := fmt.Errorf("get rabbitmq connection error %w", err)
		panic("get rabbitmq error, cause: " + err.Error())
	}
	return session
}
