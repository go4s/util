package future

import (
    "context"
    
    "github.com/go4s/util/result"
)

type (
    TaskFn[T any]          func(ctx context.Context) result.Result[T]
    Executor               interface{ Go(func() error) }
    Future[T any]          interface {
        Wait() result.Result[T]
        Done() <-chan struct{}
        Cancel()
        
        Success(T)
        Failed(error)
        Finish(result.Result[T])
    }
    
    future[T any] struct {
        root   context.Context
        cancel context.CancelFunc
        val    result.Result[T]
        notify chan struct{}
    }
)

func (a *future[T]) take() result.Result[T] {
    select {
    case <-a.notify:
        return a.val // finish
    default:
        return result.Failed[T](context.Canceled) // upstream cancel
    }
}
func (a *future[T]) Wait() result.Result[T] {
    select {
    case <-a.Done():
        return a.take()
    }
}
func (a *future[T]) Done() <-chan struct{} {
    select {
    case <-a.root.Done():
        return a.root.Done()
    default:
        return a.notify
    }
}
func (a *future[T]) Cancel()                     { a.cancel() }
func (a *future[T]) Finish(val result.Result[T]) { a.val = val; close(a.notify) }
func (a *future[T]) Success(val T)               { a.val = result.Success[T](val); close(a.notify) }
func (a *future[T]) Failed(err error)            { a.val = result.Failed[T](err); close(a.notify) }

func Submit[T any](ctx context.Context, executor Executor, fn TaskFn[T]) Future[T] {
    var t = future[T]{notify: make(chan struct{})}
    t.root, t.cancel = context.WithCancel(ctx)
    executor.Go(func() error { t.Finish(fn(t.root)); return nil })
    return &t
}

var _ Future[struct{}] = (*future[struct{}])(nil)
