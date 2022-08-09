package interfacex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilAssert(t *testing.T) {
	type user struct {
		id   string
		name string
	}

	// 已申明
	var userA user
	assert.Nil(t, userA)

	// 直接定义
	var arr []int
	// assert.Nil(t, arr[1]) // panic , 数组越界
	// assert.Nil(t, arr[0]) // panic
	assert.Nil(t, arr) // true

	// 提前申明
	arrDec := make([]int, 1)
	assert.NotNil(t, arrDec[0]) // Expected nil, but got: 0
	// assert.Nil(t, arrDec[1]) // panic

}

func TestNilAssertA(t *testing.T) {

}
