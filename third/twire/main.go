package main

import (
	"x-learn/advanced/klog/log"
	"x-learn/third/twire/twire"
)

func main() {
	app, err := twire.InitApp("User:Pass")
	if err != nil {
		log.Fatal(err)
	}
	err = app.DoThings()
	if err != nil {
		log.Fatal(err)
	}
}
