package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	address   = `kafka_dev_node:9092`
	topic     = "account-events"
	partition = 0
)

func Consumer() error {
	// 拨号
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 设置读取超时
	if err := conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal("failed to SetReadDeadline:", err)
		return err
	}

	// 批量读取
	batch := conn.ReadBatch(10e3, 1e7) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	// 批量关闭
	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	// 连接关闭
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

	return nil
}
