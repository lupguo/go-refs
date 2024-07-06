package generics

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func TestGMin(t *testing.T) {
	// 泛型调用
	x := GMin[int](2, 3)
	y := GMin[float32](2.0, 3)
	z := GMin(2.0, 3)

	t.Log(x, y, z)
}

// 拼接成SQL中的(1,2..)或者("a","b",..)形式
func Join[S ~[]E, E any](sdata S) string {
	str := "("
	for i, v := range sdata {
		if i == 0 {
			str += fmt.Sprintf(`"%v"`, v)
			continue
		}
		str += fmt.Sprintf(`,"%v"`, v)
	}
	str += ")"

	return str
}

func TestJoin(t *testing.T) {
	nums := []int{1, 2, 3}
	chars := []string{"a", "b", "c"}

	t.Log(Join(nums))
	t.Log(Join(chars))
}
