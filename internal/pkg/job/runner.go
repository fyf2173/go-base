package job

import (
	"context"
	"fmt"
)

type Job struct {
	Cmd  string
	Args []string
}

type Handler func(context.Context, []string) error

type Runner struct {
	Job
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

func (r *Runner) Exec(ctx context.Context, name string, args []string) error {
	if name == "" {
		return fmt.Errorf("unsupported handler")
	}
	f := r.handlers[name]
	return f(ctx, args)
}

func (r *Runner) ExecJob(ctx context.Context, job Job) error {
	if job.Cmd == "" {
		return fmt.Errorf("unsupported handler")
	}
	f := r.handlers[job.Cmd]
	return f(ctx, job.Args)
}
