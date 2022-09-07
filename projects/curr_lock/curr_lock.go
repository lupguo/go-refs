package curr_lock

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

var lock bool

func DoThingsV1() error {
	// defer unlock
	defer func() {
		lock = false
		log.Printf("defer unlock and set lock=%v", lock)
	}()
	// check lock after defer
	if lock {
		return errors.New("has lock, try again")
	}

	// mock handle
	lock = true
	log.Printf("do business handle")
	time.Sleep(time.Second)
	return nil
}

func DoThingsV2() error {
	// check lock first
	if lock {
		return errors.New("has lock, try again")
	}
	// defer after check lock
	defer func() {
		log.Printf("defer unlock")
		lock = false
	}()

	// mock handle
	lock = true
	log.Printf("do business handle")
	time.Sleep(time.Second)
	return nil
}