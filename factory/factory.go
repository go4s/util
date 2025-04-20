package factory

import (
    "sync"
)

type (
    Initializer[T any] func(T) error
    Factory[T any]     struct {
        once       *sync.Once
        onRegister Initializer[T]
        Instance   T
    }
)

func New[T any](initialize Initializer[T], once *sync.Once) *Factory[T] {
    return &Factory[T]{onRegister: initialize, once: once}
}
func (fact *Factory[T]) Register(inst T) (err error) {
    if fact.Instance = inst; fact.once != nil {
        fact.once.Do(func() { err = fact.onRegister(inst) })
    }
    return
}

func (fact *Factory[T]) MustRegister(inst T) {
    if err := fact.Register(inst); err != nil {
        panic(err)
    }
}
