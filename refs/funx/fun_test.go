package funx

import "testing"

func TestPointVal(t *testing.T) {
	v1, v2, v3, v4 := "hello", "hello", "hello", "hello"
	Op(v1, &v2, v3, v4)
	t.Logf("v1=%s, v2=%s, v3=%s, v4=%s", v1, v2, v3, v4)

	v1, v2, v3, v4 = "hello", "hello", "hello", "hello"
	Op(v1, &v2, &v3, v4)
	t.Logf("v1=%s, v2=%s, v3=%s, v4=%s", v1, v2, v3, v4)
}

func Op(v1 string, v2 *string, v3 interface{}, v4 interface{}) {
	v1 = "world"
	*v2 = "world"
	v3 = "world"
	v4 = v3
}
