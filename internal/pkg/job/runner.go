package job

import (
	"context"
	"fmt"
)

type Handler func(context.Context) error

type Runner struct {
	handlers map[string]Handler
}

func NewRunner() *Runner {
	return &Runner{
		handlers: make(map[string]Handler),
	}
}

func (r *Runner) Add(name string, handler Handler) {
	r.handlers[name] = handler
}

func (r *Runner) Exec(name string) error {
	if name == "" {
		return fmt.Errorf("unsupported handler")
	}
	f := r.handlers[name]
	return f(context.Background())
}
