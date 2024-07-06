package slice

import (
	"testing"
)

type Stu struct {
	ID   int
	Name string
}

func TestArrIsNil(t *testing.T) {
	a := []*Stu{} // 定义了一个值为*Stu指针的切片，a != nil
	var b []*Stu  // 定义了一个值为*Stu指针的切片， a != nil

	t.Logf("a=>%+v, b=>%+v", a, b)
	t.Logf("a=>%p, b=>%p", &a, &b)
	// t.Logf("a=b?=>%+v", a==b) // 切片不支持比较
}

func TestPointerArrIsNil(t *testing.T) {

}
