package optional

type Optional[T any] struct {
    val *T
}

func (o Optional[T]) Some() T    { return *o.val }
func (o Optional[T]) None() bool { return o.val == nil }

func Some[T any](val T) Optional[T] { return Optional[T]{&val} }
func None[T any]() Optional[T]      { return Optional[T]{nil} }
