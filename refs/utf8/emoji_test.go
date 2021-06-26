package utf8

import (
	"testing"
)

func TestEmoji(t *testing.T) {
	var emojis = []interface{}{
		'😂',
		'\uFFFD',
		0x1F601,
		0x1F602,
		0x1F603,
	}
	for _, emj := range emojis {
		t.Logf("%c, %[1]T, %#[1]v,%#[1]v, %#[1]X", emj)
	}
}
