package checknil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID uint64
}

func TestCheckNil(t *testing.T) {
	var u1 *User
	var u2 User
	var u3 = User{}
	var u4 = &User{}

	assert.Nil(t, u1, "u1 should nil")
	assert.NotNil(t, u2, "u2 is zero value")
	assert.NotNil(t, u3, "u3 should not nil")
	assert.NotNil(t, u4, "u4 should not nil")

	t.Logf("u2.ID: %d", u2.ID)
}