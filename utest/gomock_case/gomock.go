package gomock_case

import (
	"log"
)

// Data 定义一个数据接口CRUD
type Data interface {
	Send(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type AnalysisData struct {
}

func (a *AnalysisData) Send(key string, value []byte) error {
	log.Printf("Exec Send: key=>%s, value=>%s", key, value)
	return nil
}

func (a *AnalysisData) Get(key string) ([]byte, error) {
	log.Printf("Exec Get: key=>%s", key)
	return nil, nil
}

// HandleMsg 业务代码，处理消息数据，需要做UT
//	1. 通过接口生成gomock代码
//	2. 编写UT测试HandleMsg函数
func HandleMsg(d Data) {
	// send
	err := d.Send("no_value", nil)
	if err != nil {
		log.Printf("d.Send(no_value) got err, %s", err)
		// return
	}
	// send
	err = d.Send("key", []byte("Hello"))
	if err != nil {
		log.Printf("d.Send(h1) got err, %s", err)
		// return
	}

	if _, err := d.Get("no_value"); err != nil {
		log.Printf("d.Get(no_value) got err, %s", err)
	}

	// get
	get, err := d.Get("h2")
	if err != nil {
		log.Printf("d.Get(h2) got err, %s", err)
		// return
	}
	log.Printf("d.Get(h2) got %s", get)
}
