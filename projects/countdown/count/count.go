package count

import "time"

// CalcLeftTime 输入时间，获取剩余时间
func CalcLeftTime(t time.Time) (left time.Duration) {
	return t.Sub(time.Now())
}
