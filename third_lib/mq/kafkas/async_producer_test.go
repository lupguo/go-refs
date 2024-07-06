package kafkas

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"github.com/Shopify/sarama"
)

func TestNewASyncProducer(t *testing.T) {
	type args struct {
		ctx        context.Context
		name       string
		brokersUrl []string
		cfg        *sarama.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *AsyncProducer
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewASyncProducer(tt.args.ctx, tt.args.name, tt.args.brokersUrl, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewASyncProducer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewASyncProducer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAsyncSendProducer(b *testing.B) {
	b.SkipNow()

	ctx := context.Background()
	name := `asyncT`
	brokersUrl := []string{"kafka_dev_node:9092"}
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			got, err := NewASyncProducer(ctx, name, brokersUrl, nil)
			if err != nil {
				b.Error("new async producer got err:", err)
				return
			}
			err = got.AsyncPushMessageToQueue(ctx, `my-topic`, []byte(`{ID:1}`))
			if err != nil {
				b.Error("send mq got err:", err)
				return
			}
		}()
	}
	wg.Wait()
}
