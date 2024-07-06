package main

import (
	"context"

	"x-learn/mq/kafkahttpd/kconsumer"
)

var (
	brokersUrl = []string{
		"kafka_dev_node:9092",
		// "kafka_dev_node:9092",
	}
	ctx     = context.Background()
	topic   = "comments"
	groupID = "my-group"
)

func main() {
	consumer := kconsumer.NewKConsumer(brokersUrl)
	err := consumer.ConsumeMessageV2(ctx, topic, groupID)
	if err != nil {
		panic(err)
	}
}
