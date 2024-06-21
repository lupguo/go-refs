// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package xstruct

import (
	"context"
	"github.com/google/wire"
	"io"
)

// Injectors from wire.go:

// 新的注入器实现
func InitXStructGreet(msgs []string, writer io.Writer) (*Greeter, error) {
	contextContext := context.Background()
	options := &Options{
		Messages: msgs,
		Writer:   writer,
	}
	greeter, err := NewGreeter(contextContext, options)
	if err != nil {
		return nil, err
	}
	return greeter, nil
}

// wire.go:

// 注入器集合
var GreeterSet = wire.NewSet(context.Background, wire.Struct(new(Options), "*"), NewGreeter)