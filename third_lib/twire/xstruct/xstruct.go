package xstruct

import (
	"context"
	"io"
)

type Greeter struct {
}

type Options struct {
	// Messages is the set of recommended greetings.
	Messages []string
	// Writer is the location to send greetings. nil goes to stdout.
	Writer io.Writer
}

func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	// ...

	return nil, nil
}
