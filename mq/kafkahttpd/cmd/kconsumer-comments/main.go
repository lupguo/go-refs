package main

import (
	"context"

	"x-learn/mq/kafkahttpd/kconsumer"
)

var brokersUrl = []string{
	"kafka_dev_node:9092",
	"kafka_dev_node:9092",
}

func main() {
	ctx := context.Background()
	topic := "comments"
	// group consumer: 按组消费
	// 	说明:
	// 	1、当topic有3个分区，生产者会均衡的发送消息到各个partition分区，组消费者会均衡的从kafka的分区取出消息消费，若消费者数量
	// 	超过分区个数则不会被收到消息消费（比如当有4个消费者，那么有一个消费者不会收到消息）
	// 	2、当重启了消费者进程，其他消费者进程会断开，类使用重新分配
	// 	3、可以扩充分区个数，和增加消费者的数量，增加系统吞吐
	//
	gconsumer := kconsumer.NewGroupConsumer(brokersUrl)
	groupID := "my-group"
	err := gconsumer.ConsumeMessageByGroup(ctx, groupID, topic)
	if err != nil {
		panic(err)
	}

	// partition consumer
	// var partition int32 = 0
	// pconsumer := kconsumer.NewPartitionConsumer(brokersUrl)
	// err := pconsumer.ConsumeMessage(ctx, partition, topic)
	// if err != nil {
	// 	panic(err)
	// }
}
