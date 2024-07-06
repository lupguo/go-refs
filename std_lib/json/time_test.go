package json

import (
	"encoding/json"
	"testing"
	"time"
)

type CalleeOpts struct {
	Namespace   string        `json:"namespace"`
	ServiceName string        `json:"service_name"`
	Timeout     time.Duration `json:"timeout"`
}

func TestCalleeOpts(t *testing.T) {
	c := CalleeOpts{
		Namespace:   "Test",
		ServiceName: "trpc.ketang.account_bind.AccountBind",
		Timeout:     3 * time.Second,
	}
	marshal, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", marshal)
}

func TestCallOptsUnmarsh(t *testing.T) {
	s := `{"namespace":"Test","service_name":"trpc.ketang.account_bind.AccountBind","timeout":3000000000}`
	var callee CalleeOpts
	err := json.Unmarshal([]byte(s), &callee)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callee)

}
