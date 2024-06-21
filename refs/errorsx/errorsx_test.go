package errorsx

import (
	"testing"

	"github.com/pkg/errors"
)

func TestShowErr(t *testing.T) {
	var (
		err1 = errors.New("error 1")
		err2 = errors.New("error 2")
	)

	var err error
	err = err1
	err = err2
	t.Logf("%v", errors.Is(err, err1))
	t.Logf("%v", errors.Is(err, err2))
}
