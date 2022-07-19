package base16

import (
	"bytes"
	"testing"
)

func Base16Encode(strBin string) string {
	var encodeBytes bytes.Buffer
	sHex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	l := len(strBin)
	for i := 0; i < l; i++ {
		encodeBytes.WriteString(sHex[uint8(strBin[i])>>4])
		encodeBytes.WriteString(sHex[uint8(strBin[i])&15])
	}
	return encodeBytes.String()
}

func TestBase16(t *testing.T) {
	t.Logf(Base16Encode("hello"))
}
