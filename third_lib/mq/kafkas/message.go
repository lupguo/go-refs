package kafkas

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// SprintfMessage 格式返回消息内容，方便做Log打印
func SprintfMessage(msg *sarama.ConsumerMessage) string {
	return fmt.Sprintf(
		"Message Time(%v) | Partition(%d) | Topic(%s) | Message(%s) | Offset(%d)",
		msg.Timestamp,
		msg.Partition,
		msg.Topic,
		msg.Value,
		msg.Offset,
	)
}
