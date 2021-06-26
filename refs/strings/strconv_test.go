package strings

import (
	"strconv"
	"testing"
)

func TestFormatUint(t *testing.T) {
	var uid uint64 = 144115232311496011
	t.Logf("%s", strconv.FormatUint(uid, 10))
}
