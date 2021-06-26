package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	// Bool判断
	assert.True(t, true, "True is true!")
	assert.False(t, false, "False is false!")

	// Equal判断
	assert.Equal(t, 123, 123, "they should be equal")
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var object map[string]int
	assert.Nil(t, object, "object should be nil value")
	object = make(map[string]int)
	assert.NotNil(t, object, "now, object should not be nil")
}

func TestSomething1(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(123, 123, "they should be equal")
	assertions.NotEqual(123, 456, "they should not be equal")
}
