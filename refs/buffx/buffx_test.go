package buffx

import (
	"strings"
	"testing"

	// "go.uber.org/zap/buffer"

	"go.uber.org/zap/buffer"
)

func TestAdd(t *testing.T) {
	s := strAdd()

	t.Logf("%s", s)
}

func strAdd() string {
	var s string
	for i := 0; i < 10000; i++ {
		s += "a"
	}
	return s
}
func TestGenBuffx(t *testing.T) {
	b := zapBufferAdd()

	t.Logf("%s", b.String())
}

func zapBufferAdd() *buffer.Buffer {
	b := &buffer.Buffer{}

	for i := 0; i < 10000; i++ {
		b.AppendInt(int64(i))
	}
	return b
}

func TestStringBuilder(t *testing.T) {
	sb := stringsBuilderAdd()

	t.Logf("%s", sb.String())
}

func stringsBuilderAdd() *strings.Builder {
	sb := &strings.Builder{}
	for i := 0; i < 10000; i++ {
		sb.WriteString("s")
	}
	return sb
}

func BenchmarkStrAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAdd()
	}
}

func BenchmarkZapBufferAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zapBufferAdd()
	}
}
func BenchmarkStringsBuilderAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsBuilderAdd()
	}
}
