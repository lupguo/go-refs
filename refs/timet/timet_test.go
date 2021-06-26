package timet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeNil(t *testing.T) {
	var t1 time.Time
	var t2 time.Timer
	var t3 time.Duration
	var t4 time.Ticker
	var t5 time.Location
	t.Logf("time.Time => %#v, default value => %v, pointer => %0p", t1, t1 == time.Time{}, &t1)
	t.Logf("time.Timer => %#v", t2)
	t.Logf("time.Duration => %#v", t3)
	t.Logf("time.Ticker => %#v", t4)
	t.Logf("time.Location => %#v", t5)
}

func TestUnixTime(t *testing.T) {
	now := time.Now()
	t.Logf("now.Uinx()     => %d", now.Unix())
	t.Logf("now.Uinx()x1e3 => %d", now.Unix()*1e3)
	t.Logf("now.Uinx()x1e6=> %d", now.Unix()*1e6)
	t.Logf("now.Uinx()x1e9=> %d", now.Unix()*1e9)
	t.Logf("now.UnixNano() => %d", now.UnixNano())
}

func TestUnixToStrTime(t *testing.T) {
	tlong := int64(1599553509352)
	that := time.Unix(tlong/1e3, 0)
	tf := that.Format("2006/01/02 15:04:05")
	t.Logf("that time => %s", tf)
}

func TestRoundTime(t *testing.T) {
	ts := []string{
		"760h",
		"3h",
		"5m",
		"32s",
	}
	for _, s := range ts {
		duration, err := time.ParseDuration(s)
		assert.Nil(t, err)
		t.Logf("dur:%s, round dur:%s", duration, duration.Round(2*time.Second))
	}
}

func TestName(t *testing.T) {

}
