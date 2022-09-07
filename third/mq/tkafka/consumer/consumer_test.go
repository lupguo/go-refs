package consumer

import (
	"log"
	"testing"
)

func TestConsumer(t *testing.T) {
	// 消费
	err := Consumer()
	if err != nil {
		log.Fatalln("consumer got err:", err)
		return
	}
}
