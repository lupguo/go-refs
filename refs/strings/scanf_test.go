package strings

import (
	"fmt"
	"testing"
)

func TestSscanf(t *testing.T) {
	content := `@机构赤字小助手 查询机构赤字明细123`
	var cmd, valstr string
	_, err := fmt.Sscanf(content, `@机构赤字小助手 %s%s`, &cmd, &valstr)

	t.Logf("err=%v, cmd=%v, valstr=%v", err, cmd, valstr)
}

func TestSscanCase2(t *testing.T) {
	tt := []string{
		"",
		"var1",
		"var1 var2",
		"var1   var2",
		"var1var2",
		"12",
		"12 34",
	}

	for _, s := range tt {
		var cmd, valstr string
		n, err := fmt.Sscanf(s, "%s%s", &cmd, &valstr)
		t.Logf("n=%v, err=%v, cmd=%v, valstr=%v", n, err, cmd, valstr)
	}
}
