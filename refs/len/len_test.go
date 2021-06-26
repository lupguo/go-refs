package len

import (
	"bytes"
	"testing"
)

const (
	// 表示QAPP协议的开始符号
	QAPP_STX_C = 0x2
	// 表示QAPP协议的终止符号
	QAPP_ETX_C = 0x3
)

func TestLen(t *testing.T) {
	var s,e bytes.Buffer
	s.WriteByte(QAPP_STX_C)
	t.Logf("len(b)=%d", s.Len())
	e.WriteByte(QAPP_ETX_C)
	t.Logf("len(b)=%d", e.Len())

	t.Logf("int 0x2=%d, 0x3=%d", QAPP_STX_C,QAPP_ETX_C)
}
