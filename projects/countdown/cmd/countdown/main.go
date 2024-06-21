package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"x-learn/projects/countdown/count"
)

var d = flag.String("date", "2021/01/01", "input calc date")

func main() {
	flag.Parse()
	dt, err := time.Parse("2006/01/02", *d)
	if err != nil {
		fmt.Printf("incorrect date format, %v\n", err)
		return
	}
	leftTime := count.CalcLeftTime(dt)
	hs := strconv.FormatFloat(leftTime.Minutes()/(24*60), 'f', 0, 64)
	fmt.Printf("only %s days left until %s\n", hs, *d)
}
