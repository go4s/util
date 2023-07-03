package iterator

import "github.com/go4s/util/optional"

type (
    array[T any] struct {
        items []T
        idx   uint
    }
)

func (r *array[T]) Next() optional.Optional[T] {
    if m, c := uint(len(r.items)), r.idx; c < m {
        r.idx += 1
        return optional.Some[T](r.items[c])
    }
    return optional.None[T]()
}
func (r *array[T]) At() uint { return r.idx }

func New[T any](items ...T) Iterator[T] {
    iter := array[T]{[]T{}, 0}
    iter.items = append(iter.items, items...)
    return &iter
}
