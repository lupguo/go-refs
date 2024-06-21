package greet

import (
	"fmt"
)

// Message 消息
type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

// Speaker 迎宾
type Speaker interface {
	Greet() Message
}

// HumanGreeter 迎宾员
type HumanGreeter struct {
	Message Message // <- adding a Message field
}

func NewHumanGreeter(m Message) HumanGreeter {
	return HumanGreeter{Message: m}
}

func (g HumanGreeter) Greet() Message {
	return g.Message
}

// RobotGreeter 机器人迎宾
type RobotGreeter struct {
	Message Message
	ID      string `wire:"-"`
}

func NewRobotGreeter(message Message, ID string) *RobotGreeter {
	return &RobotGreeter{
		Message: message,
		ID:      ID,
	}
}

func (r *RobotGreeter) Greet() Message {
	return Message(fmt.Sprintf("I'am a Robot[%v]: Welcome %v", r.ID, r.Message))
}

// Event 时间
type Event struct {
	Speaker Speaker // <- adding a Speaker field
}

func NewEvent(spk Speaker) Event {
	return Event{spk}
}

func (e Event) Start() {
	msg := e.Speaker.Greet()
	fmt.Println(msg)
}
