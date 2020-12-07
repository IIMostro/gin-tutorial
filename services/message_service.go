package services

import (
	"ilmostro.org/gin-tutorial/configuration"
	"log"
)

func init() {
	log.Printf("rabbit connection aleardy!,start consumer")
	consumer("insert-user-queue")
}

func consumer(queue string) {
	consume, err := configuration.GetRabbitConnection().Consume(
		queue,
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
		//防止阻塞主线程。
		for msg := range consume {
			log.Printf(" consumer message: %s", msg.Body)
		}
	}()
}
