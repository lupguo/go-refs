package pointerx

import (
	"testing"
)

type User struct {
	ID int
}

func SetMagicCodeForApp(u *User) {
	u.ID = 100
}

func TestUpdateUser(t *testing.T) {
	u := &User{}
	t.Logf("user %v", u)
}
