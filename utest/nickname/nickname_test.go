package nickname

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func genNickName(s string) string {
	newNickName := ""
	for _, s := range s {
		if s >= '0' && s <= '9' {
			newNickName += "*"
		} else {
			newNickName += string(s)
		}
	}
	return newNickName
}

func TestNickname(t *testing.T) {
	ts := []struct {
		name string
		want string
	}{
		{"185****5717", "185****5717"},
	}
	for _, s := range ts {
		assert.Equal(t, s.want, genNickName(s.name))
	}
}
