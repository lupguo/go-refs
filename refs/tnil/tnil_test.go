package tnil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type U struct {
	Id string
}

func TestNil(t *testing.T) {
	var u1 U
	assert.Equal(t, u1, U{})
	u2 := U{}
	assert.NotNil(t, u2)
	//
	su1 := []U{}
	assert.NotNil(t, su1)

	var su2 []U
	assert.Nil(t, su2)
}
