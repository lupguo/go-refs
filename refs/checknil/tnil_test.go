package checknil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceNil(t *testing.T) {
	// 切片
	sliceA := []User{} // 申明了User类型切片，已初始化了内存空间(长度为0)，sliceA指向该区域，值不为nil
	t.Logf("sliceA=>%T, p=>%[1]p, v=%[1]v", sliceA)
	assert.NotNil(t, sliceA)

	// 切片
	var sliceB []User // 申明了User结构体切片，但未做初始化，sliceB指向该区域，所以SliceB的值为nil，var slice []Type 这种为nil, sliceA=>[]tnil.User, p=>0x0, v=[]
	t.Logf("sliceB=>%T, p=>%[1]p, v=%[1]v", sliceB)
	assert.Nil(t, sliceB)

	var sliceC []*User // 同上，申明了存储指针*User的切片，但为做初始化，SliceC指向该区域，所以SliceC为nil
	t.Logf("sliceC=>%T, p=>%[1]p, v=%[1]v", sliceC)
	assert.Nil(t, sliceC)

	var sliceD = []*User{} // 申明了指针*User的切片，已做了初始化，SliceD指向该区域，所以sliceD不为nil
	t.Logf("sliceD=>%T, p=>%[1]p, v=%[1]v", sliceD)
	assert.NotNil(t, sliceD)

	// sliceA != sliceD，一个切片的值类型为*User，一个是User，不相等
	assert.NotEqual(t, sliceA, sliceD)
}

// p=>%[1]p 不能应用于非指针
// var StructA User, 并非nil，即初始化了一个空结构体，可以通过指针取值
func TestStructNil(t *testing.T) {
	var StructA User // 已申明内存空间了, &StructA不为空
	t.Logf("StructA=>%T, v=%[1]v， p=>%p", StructA, &StructA)
	assert.NotNil(t, StructA) //

	// if StructA != User(nil) {} 编译错误
	assert.NotNil(t, &StructA) // StructA已申明过了，所以肯定不会为nil

	var StructB *User // 申明了一个指向User类型的指针，但其值为nil
	t.Logf("StructB=>%T, v=%[1]v， p=>%p", StructB, &StructB)
	assert.Nil(t, StructB) //

	StructC := User{} // 初始化User结构体，并将StructC指针指向该结构体， 所以StructC不为空
	t.Logf("StructC=>%T, v=%[1]v， p=>%p", StructC, &StructC)
	assert.NotNil(t, StructC) //

	var StructD = &User{} // 初始化User结构体，并将StructC指针指向该结构体， 所以StructC不为空(通常会适用StructC类型方式简短定义)
	t.Logf("StructD=>%T, v=%[1]v， p=>%p", StructD, &StructD)
	assert.NotNil(t, StructD) //
}
