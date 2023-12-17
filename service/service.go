package service

import (
    "context"
    "sync"
    
    "github.com/go4s/util/future"
)

type (
    StarterFn func(context.Context, future.Executor) error
    Starter   interface {
        Start(context.Context, future.Executor) error
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
func Start(ctx context.Context, executor future.Executor) (err error) {
    for _, svc := range services {
        if err = svc.Start(ctx, executor); err != nil {
            return
        }
    }
    return
}
func (fn StarterFn) Start(ctx context.Context, executor future.Executor) error {
    return fn(ctx, executor)
}
