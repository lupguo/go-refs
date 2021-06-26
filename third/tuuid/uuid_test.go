package tuuid

import (
	"strconv"
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUUIDV4(t *testing.T) {
	// Creating UUID Version 4
	u1, u2 := uuid.NewV4(), uuid.NewV4()
	assert.Equalf(t, u1.String(), u2.String(), "u1:%s\nu2:%s", u1.String(),u2.String())

	// // Parsing UUID from string input
	// u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	// if err != nil {
	// 	t.Logf("Something gone wrong: %s", err)
	// }
	// t.Logf("Successfully parsed: %s", u2)
}

func TestUUIDV5(t *testing.T) {
	u1 := uuid.NewV5(uuid.NamespaceURL, "www.example.com")
	u2 := uuid.NewV5(uuid.NamespaceURL, "www.example.com")

	// assert.Equal(t, u1, u2, "should equal")
	assert.Equal(t, u1.String(), u2.String() ,"should not equal")
	assert.NotEqualf(t, u1.String(),u2.String(), "u1:%s\nu2:%s", u1.String(), u2.String())
}

func TestUUIDV5Rand(t *testing.T) {
	str1 := strconv.FormatInt(time.Now().UnixNano(), 10)
	u1 := uuid.NewV5(uuid.FromStringOrNil(str1), "hello")
	str2 := strconv.FormatInt(time.Now().UnixNano(), 10)
	u2 := uuid.NewV5(uuid.FromStringOrNil(str2), "hello")

	assert.Equal(t, u1.String(), u2.String() ,"should not equal")
	assert.NotEqualf(t, u1.String(), u2.String() ,"should not equal")
}


