package timex

import (
	"testing"
	"time"
)

func Test2038(t *testing.T) {
	dtime, err := time.Parse("2006/01/02 15:04:05", "2099/01/02 15:04:05")
	if err != nil {
		return
	}
	t.Logf("%v", dtime)

	dtimeStamp := dtime.Unix()
	t.Logf("%v", dtime.Unix())

	// 2099年时间戳转日期
	dtime = time.Unix(dtimeStamp, 0)
	t.Logf("%v", dtime)
}
