package switchcase

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	vks := []interface{}{
		6,
		6.66,
		true,
	}
	for _, vk := range vks {
		switch vv := vk.(type) {
		case int,float64:
			fmt.Println("case number", vv)
		case bool:
			fmt.Println("bool values", vv)
		}
	}
}
