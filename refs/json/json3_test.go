package json

import (
	"encoding/json"
	"testing"
)

func TestJson3(t *testing.T) {
	customXgData := map[string]interface{}{
		"channel":  "xg",
		"push_id":  "0",
		"pushtype": "4", // 代表运营类型，目前app只接入这个类型
		"data":     (map[string]interface{}{"schema": "https://ke.qq.com"}),
	}

	marshal, err := json.Marshal(customXgData)
	if err != nil {
		t.Errorf("marshal got err: %s", err)
	}

	t.Logf("%s", marshal)
}

func BenchmarkMarshalUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
	}
}
