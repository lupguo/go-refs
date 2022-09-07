package producer

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func ProduceMessage02() {
	producer, err := sarama.NewAsyncProducer([]string{"kafka_dev_node:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, producerErrors int
ProducerLoop:
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{
			Topic: "my_topic",
			Key:   nil,
			Value: sarama.StringEncoder("testing 123"),
		}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			producerErrors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, producerErrors)
}
