package pointer_set

import (
	"testing"
)

type User struct {
	ID uint64
}

func SetMagicCodeForApp(u *User) {
	u.ID = 100
}

func TestUpdateUser(t *testing.T) {
	u := &User{}
	t.Logf("user %v", u)
}
