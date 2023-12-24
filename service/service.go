package service

import (
    "context"
    "sync"
)

type (
    Executor interface {
        Go(func() error)
    }
    StarterFn func(context.Context, Executor) error
    Starter   interface {
        Start(context.Context, Executor) error
    }
)

var (
    services    []Starter
    initializer = sync.Once{}
)

func init() {
    initializer.Do(initialize)
}

func initialize() { services = []Starter{} }

func Register(svc Starter) { services = append(services, svc) }
func Start(ctx context.Context, executor Executor) (err error) {
    for _, svc := range services {
        if err = svc.Start(ctx, executor); err != nil {
            return
        }
    }
    return
}
func (fn StarterFn) Start(ctx context.Context, executor Executor) error {
    return fn(ctx, executor)
}
