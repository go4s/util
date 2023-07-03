package iterator

import "github.com/go4s/util/optional"

type (
    Iterator[T any] interface {
        Next() optional.Optional[T]
        At() uint
    }
    RingIterator[T any] interface {
        Iterator[T]
        Remove(uint) optional.Optional[T]
        Empty() bool
    }
)
