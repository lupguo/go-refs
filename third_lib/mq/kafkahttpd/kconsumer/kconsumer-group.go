package kconsumer

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type GroupConsumer struct {
	BrokersUrl []string
}

func NewGroupConsumer(brokersUrl []string) *GroupConsumer {
	return &GroupConsumer{
		BrokersUrl: brokersUrl,
	}
}

func (gc *GroupConsumer) ConnectConsumerGroups(groupID string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	gconsumer, err := sarama.NewConsumerGroup(gc.BrokersUrl, groupID, config)
	if err != nil {
		return nil, err
	}
	return gconsumer, nil
}

//
// func (gc *GroupConsumer) CreateTopic(topic string, parition int) {
// 	b := sarama.NewBroker("d")
// 	var topicReq = &sarama.CreateTopicsRequest{
// 		NowVersion:      0,
// 		TopicDetails: nil,
// 		Timeout:      0,
// 		ValidateOnly: false,
// 	}
// 	_, err := b.CreateTopics(topicReq)
// 	if err != nil {
// 		return
// 	}
//
// 	var paritionReq *sarama.CreatePartitionsRequest
// 	_, err = b.CreatePartitions(paritionReq)
// 	if err != nil {
// 		return
// 	}
// }

func (gc *GroupConsumer) ConsumeMessageByGroup(ctx context.Context, groupID string, topic string) error {
	consumer, err := gc.ConnectConsumerGroups(groupID)
	if err != nil {
		return err
	}

	// Track errors
	go func() {
		for err := range consumer.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	err = consumer.Consume(ctx, []string{topic}, &WorkHandler{})
	if err != nil {
		return err
	}

	return nil
}

type WorkHandler struct{}

func (h *WorkHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *WorkHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *WorkHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message Time(%v)| Partition(%d) | Topic(%s) | Message(%s) | Offset(%d)", msg.Timestamp, msg.Partition, msg.Topic, msg.Value, msg.Offset)
		// sess.MarkMessage(msg, "")
	}
	return nil
}
