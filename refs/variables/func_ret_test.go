package variables

import (
	"fmt"
	"testing"
)

//
func fnX(v int) (ret []int) {
	fmt.Printf("%+v, %[1]T", ret)
	return ret
}

func TestGetWhiteConcatAidSet(t *testing.T) {
	v := fnX(1)
	fmt.Printf("%+v, %[1]T", v)
	t.Logf("%+v", v)
}
