package main

import (
	"fmt"

	"x-learn/third/gocobra/cmd"
)

func main() {

	// 执行命令
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
