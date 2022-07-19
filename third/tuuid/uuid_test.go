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
	assert.NotEqual(t, u1.String(), u2.String())

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
	assert.Equal(t, u1.String(), u2.String())
}

func TestUUIDV5Rand(t *testing.T) {
	str1 := strconv.FormatInt(time.Now().UnixNano(), 10)
	uuidFrom1 := uuid.FromStringOrNil(str1)
	u1 := uuid.NewV5(uuidFrom1, "hello")

	str2 := strconv.FormatInt(time.Now().UnixNano(), 10)
	uuidFrom2 := uuid.FromStringOrNil(str2)
	u2 := uuid.NewV5(uuidFrom2, "hello")

	t.Logf("str1=>%s, str2=>%s", str1, str2)
	t.Logf("uuid.FromStringOrNil str1=>%s, uuid.FromStringOrNil str2=>%s", uuidFrom1, uuidFrom2)
	assert.Equal(t, u1.String(), u2.String())

	u3, _ := uuid.FromBytes([]byte(str1))
	u4, _ := uuid.FromBytes([]byte(str2))
	t.Logf("u3=>%s, u4=>%s", u3, u4)
	assert.NotEqual(t, u3, u4)
}
