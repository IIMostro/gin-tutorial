package services

import (
	"log"
)

func init() {
	log.Printf("rabbit connection aleardy!,start consumer")
	consume, err := channel.Consume(
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
