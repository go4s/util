package ring

import (
    "github.com/go4s/util/iterator"
    "github.com/go4s/util/optional"
)

type (
    ring[T any] struct {
        items []T
        idx   uint
    }
)

func (r *ring[T]) Next() optional.Optional[T] { r.advance(); return optional.Some[T](r.items[r.idx]) }
func (r *ring[T]) advance()                   { r.idx = (r.idx + 1) % uint(len(r.items)) }
func (r *ring[T]) At() uint                   { return r.idx }
func (r *ring[T]) Empty() bool                { return len(r.items) == 0 }
func (r *ring[T]) Remove(idx uint) optional.Optional[T] {
    if idx >= uint(len(r.items)) {
        return optional.None[T]()
    }
    var item T
    item, r.items = r.items[idx], append(r.items[:idx], r.items[idx+1:]...)
    if idx > r.idx {
        r.advance()
    }
    return optional.Some(item)
}

func New[T any](items ...T) iterator.RingIterator[T] {
    iter := ring[T]{[]T{}, 0}
    iter.items = append(iter.items, items...)
    return &iter
}
