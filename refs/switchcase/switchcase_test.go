package switchcase

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSwitch(t *testing.T) {
	vks := []interface{}{
		6,
		6.66,
		true,
	}
	for _, vk := range vks {
		switch vv := vk.(type) {
		case int, float64:
			fmt.Println("case number", vv)
		case bool:
			fmt.Println("bool values", vv)
		}
	}
}

func TestSwitchMuchCase(t *testing.T) {
	isOddFn := func(n int32) bool {
		switch n {
		case 1, 3: // 1,3
			fallthrough
		case 5: // 5
			return true
		}
		return false
	}

	tests := []struct {
		name string
		n    int32
		want bool
	}{
		{"1", 1, true},
		{"2", 2, false},
		{"3", 3, true},
	}
	for _, tt := range tests {
		rs := isOddFn(tt.n)
		err := errors.New(fmt.Sprintf("isOddFn(%d) got %t, but want %t", tt.n, rs, tt.want))
		assert.Equal(t, tt.want, rs, err)
	}
}
