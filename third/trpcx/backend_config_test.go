package apollo_conf

import (
	"encoding/json"
	"testing"

	"git.code.oa.com/trpc-go/trpc-go/client"
)

var backendService = `{
	 "ServiceName": "go_edu_tim_push",
	 "Namespace": "Test",
	 "Target": "polars://xxx",
	 "Network": "tcp",
	 "Timeout": 500,
	 "Protocol": "qapp"
}`

func TestGenerateBackend(t *testing.T) {
	callee := &client.BackendConfig{}
	b, err := json.MarshalIndent(callee, " ", "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", b)
}

func TestUnmarshalBackend(t *testing.T) {
	var cfg client.BackendConfig
	err := json.Unmarshal([]byte(backendService), &cfg)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", cfg)
}
