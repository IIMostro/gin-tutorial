package services

import (
	"ilmostro.org/gin-tutorial/configuration"
	"log"
)

func init() {
	log.Printf("rabbit connection aleardy!,start consumer")
	consume, err := configuration.GetRabbitConnection().Consume(
		"insert-user-queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return
	}
	go func() {
		for msg := range consume {
			log.Printf(" consumer message: %s", msg.Body)
		}
	}()
}
