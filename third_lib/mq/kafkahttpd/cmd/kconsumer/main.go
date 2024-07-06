package main

import (
	"context"
	"log"

	"github.com/spf13/pflag"
	"x-learn/mq/kafkahttpd/kconsumer"
)

var (
	address string
	topic   string
	groupID string
)

func main() {
	pflag.StringVarP(&address, "address", "a", "11.186.10.169:9092", "kafka address")
	pflag.StringVarP(&topic, "topic", "t", "account_event", "kafka topic")
	pflag.StringVarP(&groupID, "group_id", "g", "console-group","consumer group id")
	pflag.Parse()

	// consume topic from kafka
	ctx := context.Background()
	gconsumer := kconsumer.NewGroupConsumer([]string{address})

	err := gconsumer.ConsumeMessageByGroup(ctx, groupID, topic)
	if err != nil {
		log.Println(err)
	}
}
