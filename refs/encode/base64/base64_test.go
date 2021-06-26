package base64

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	input := []byte("😂")
	output := bytes.NewBuffer(make([]byte, 10))
	encoder := base64.NewEncoder(base64.StdEncoding, output)
	if _, err := encoder.Write(input); err != nil {
		return
	}
	if err := encoder.Close(); err != nil {
		return
	}
	t.Logf("%s", output)
}

func TestStdEncoding(t *testing.T) {
	input := []byte("😂")
	output := make([]byte, 100)
	base64.StdEncoding.Encode(output, input)
	t.Logf("%s", output)
	t.Logf("len(input)=%d, len(output)=%d", len(input), len(output))

	// nopadding
	base64.StdEncoding.WithPadding(base64.NoPadding)
	base64.StdEncoding.Encode(output, input)
	t.Logf("%s", output)
	t.Logf("len(input)=%d, len(output)=%d", len(input), len(output))
}

func TestDecoding(t *testing.T) {
	src := []byte(`8J+Ygg==`)
	dst := make([]byte, 10)
	base64.StdEncoding.Decode(dst, src)
	t.Logf("%s", dst)
}
