package files

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	// 创建临时文件
	fileTmp, err := ioutil.TempFile("/tmp", "_pattern")
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", *fileTmp)
	stat, err := os.Stat(fileTmp.Name())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", stat)
}
