package snowflakex

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/rand"
)

// 启动 HTTP 服务器
func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

type cache struct {
	data sync.Map // 缓存数据存在并发map中
	mu   sync.RWMutex
	ch   chan bool
}

var c cache

func init() {
	if c.ch == nil {
		c.ch = make(chan bool) // 阻塞chan
	}

	//
	go func() {
		for {
			select {
			case <-time.Tick(3 * time.Second):
				c.ch <- true
			}
		}
	}()
}

func isCrash(id string) bool {
	select {
	case <-c.ch:
		c.mu.Lock()
		log.Println("3 second after, get write lock, now to reset c.sm")
		defer c.mu.Unlock()
		c.data = sync.Map{}
		return isCrash(id)
	default:
		// sync map加载或存在指定的值
		c.mu.RLock()
		defer c.mu.RUnlock()
		if _, ok := c.data.LoadOrStore(id, id); ok {
			return true
		}
		return false
	}
}

// GenSnowflakeOrderID 22位=14位日期格式+4位机器IP+随机4位
func GenSnowflakeOrderID(maxTimes int) (string, error) {
	for i := 0; i < maxTimes; i++ {
		id := fmt.Sprintf("%s%04d%04d", time.Now().Format("20060102150405"), IPToNumber(), rand.Int31n(1e4))
		if !isCrash(id) {
			return id, nil
		}
	}
	return "", errors.New("error max times")
}

// 生成指定数量ID、检测是否冲突
func TestGenSnowflakeID(t *testing.T) {
	sema := make(chan struct{}, 100)
	sign := make(chan bool)
	for i := 0; i < 1e5; i++ {
		sema <- struct{}{}
		go func(sign chan bool) {
			if _, ok := <-sign; ok {
				return
			}
			id, err := GenSnowflakeOrderID(3)
			if err != nil {
				t.Errorf("got err: %v", err)
				return
			}
			t.Logf("got snow flake id %v", id)
			<-sema
		}(sign)
	}
}

func BenchmarkGenSnowflakeID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenSnowflakeOrderID(10)
	}
}
