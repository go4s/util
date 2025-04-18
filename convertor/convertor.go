package convertor

type (
	From[T any] interface {
		From(T) error
	}
	Into[T any] interface {
		Into() (T, error)
	}
)
