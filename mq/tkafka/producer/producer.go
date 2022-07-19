package producer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap/buffer"
)

var (
	address   = `kafka_dev_node:9092`
	topic     = "account-events"
	partition = 0
)

func MsgProducer() error {
	// 拨号创建kafka连接
	kaConn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 设置写入超时时间
	if err := kaConn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal("failed to set write deadline:", err)
	}

	for i := 0; i < 1e3; i++ {
		// 写入MQ消息
		bs := buffer.Buffer{}
		bs.AppendString(fmt.Sprintf(`{id: %d, message: the "%[1]d news"}`, i))
		nbytes, err := kaConn.WriteMessages(
			kafka.Message{
				Value: bs.Bytes(),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		log.Printf("write %d bytes!", nbytes) // 11 bytes
		time.Sleep(10 * time.Microsecond)
	}

	// 关闭连接
	if err := kaConn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	return nil
}

func MsgProducerByTick() error {
	kaconn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		panic(err)
	}

	// 单一定时器生产者
	var msgID, errcnt int
	for {
		select {
		case <-time.Tick(10 * time.Millisecond):
			// set max write time
			kaconn.SetWriteDeadline(time.Now().Add(time.Second))

			// message key
			key := []byte(uuid.New().String())

			// value
			msgID++
			bs := buffer.Buffer{}
			bs.AppendString(fmt.Sprintf(`{id:%d, datetime: %s, message: the "%[1]d news"}`, msgID, time.Now().String()))

			// message build
			messsage := kafka.Message{
				Key:   key,
				Value: bs.Bytes(),
				Time:  time.Now(),
			}

			_, err := kaconn.WriteMessages(messsage)
			if err != nil {
				errcnt++
				log.Printf("write message got err: %s, errcnt: %d", err, errcnt)
				if errcnt > 2 {
					return errors.New("exceed the maximum number of errors!")
				}
			}
		}
	}
}
