package producer

import (
	"testing"
)

func TestProducerOne(t *testing.T) {
	err := MsgProducer()
	if err != nil {
		return
	}
}

func TestMsgProducerByTick(t *testing.T) {
	err := MsgProducerByTick()
	if err != nil {
		panic(err)
	}
}