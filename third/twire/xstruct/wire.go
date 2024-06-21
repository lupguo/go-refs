//go:build wireinject
// +build wireinject

package xstruct

import (
	"context"
	"io"

	"github.com/google/wire"
)

// 注入器集合
var GreeterSet = wire.NewSet(
	context.Background,
	wire.Struct(new(Options), "*"),
	NewGreeter,
)

// 新的注入器实现
func InitXStructGreet(msgs []string, writer io.Writer) (*Greeter, error) {
	wire.Build(
		GreeterSet,
	)

	return &Greeter{}, nil
}
