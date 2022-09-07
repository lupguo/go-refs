package kafkas

import (
	"encoding/json"
	"testing"
)

func TestKafkaConfig(t *testing.T) {
	s := `["10.101.201.93:9092"]`
	var brokersUrls []string
	err := json.Unmarshal([]byte(s), &brokersUrls)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", brokersUrls)
	t.Logf("%v", brokersUrls[0])
}
