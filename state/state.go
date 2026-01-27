package state

import (
	"context"
	"sync"

	"github.com/go4s/util/convertor"
)

type (
	SnapshotNotifyFn[T any] func(context.Context, T)
	Modifier[T any]         func(T) T

	SnapshotUpdater[T any] func(context.Context, ...Modifier[T]) error

	Watcher[T any] interface {
		Watch(SnapshotNotifyFn[T])
	}
	Remover[T any] interface {
		Delete(context.Context) error
		DeleteWithMutex(context.Context, sync.Locker) error
	}
	Migrator[T any] interface {
		Migrate(context.Context, string) error
		MigrateWithMutex(context.Context, sync.Locker, string) error
	}
	Updater[T any] interface {
		Update(context.Context, ...Modifier[T]) error
		UpdateWithMutex(context.Context, sync.Locker, ...Modifier[T]) error
	}
	Mutator[T any] interface {
		Set(context.Context, convertor.Into[T]) error
		SetWithMutex(context.Context, sync.Locker, convertor.Into[T]) error
	}
	Fetcher[T any] interface {
		Get(context.Context, convertor.From[T]) error
		GetWithMutex(context.Context, sync.Locker, convertor.From[T]) error
	}

	State[T any] interface {
		Fetcher[T]
		Mutator[T]
		Updater[T]
		Watcher[T]
		Remover[T]
		Migrator[T]
		Mode() int
		Key() string
	}
	DecoratorFn[T any] func(State[T])
)
