package kafkas

import (
	"context"
	"encoding/json"
	"testing"
)

func TestNewProducerComb(t *testing.T) {
	// skip for remote test
	// t.SkipNow()

	// real check
	brokerURLs := []string{
		"10.101.201.93:9092",
	}
	ctx := context.Background()
	syncProducer, err := newSyncProducer(ctx, brokerURLs, nil)
	if err != nil {
		t.Fatal(err)
	}
	asyncProducer, err := NewASyncProducer(ctx, "newAsync", brokerURLs, nil)
	if err != nil {
		t.Fatal(err)
	}

	// new producer combine
	producer := NewProducer(asyncProducer, syncProducer)

	// sync send message
	topic := "my-topic"
	user1 := map[int]string{
		1001: "clark",
	}
	message, _ := json.Marshal(user1)
	err = producer.SyncPushMessageToQueue(ctx, topic, message)
	if err != nil {
		t.Fatal(err)
	}

	// async send message
	user2 := map[int]string{
		1002: "terry",
	}
	message, _ = json.Marshal(user2)
	err = producer.AsyncPushMessageToQueue(ctx, topic, message)
	if err != nil {
		t.Fatal(err)
	}
}
