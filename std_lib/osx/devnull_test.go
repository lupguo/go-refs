package osx

import (
	"fmt"
	"os"
	"testing"
)

func TestFprintfDevNull(t *testing.T) {
	// 将标准输出重定向到/dev/null
	devNull, _ := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	defer devNull.Close()
	oldStdout := os.Stdout
	os.Stdout = devNull

	// 执行你的代码
	fmt.Println("This will be discarded")

	// 恢复标准输出
	os.Stdout = oldStdout

	// 输出到标准输出
	fmt.Println("This will be printed")
}
