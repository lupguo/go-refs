package json

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJsonMarshalTime(t *testing.T) {
	type ClientService struct {
		Namespace string
		Timeout   int
		Target    string
	}

	cli := &ClientService{
		Namespace: "Test",
		Target:    "polaris://trpc.ketang.go_edu_tim_push.GoEduTIMPushQapp",
		Timeout:   500,
	}

	v, err := json.Marshal(cli)
	if err != nil {
		t.Error(err)
	}

	t.Logf("\n%s", v)

	s := `{"Namespace":"Test","Timeout":500,"Target":"polaris://trpc.ketang.go_edu_tim_push.GoEduTIMPushQapp"}`

	var srvName *ClientService
	err = json.Unmarshal([]byte(s), &srvName)
	if err != nil {
		t.Error(err)
	}
	var k time.Duration
	k = time.Duration(srvName.Timeout) * time.Millisecond

	t.Logf("%+v, k=>%v", srvName, k)
}
