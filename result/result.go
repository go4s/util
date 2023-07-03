package result

import "encoding/json"

type (
    Result[T any] struct {
        e error
        v T
    }
)

func Success[T any](val T) Result[T]    { return Result[T]{v: val} }
func Failed[T any](err error) Result[T] { return Result[T]{e: err} }

func (r Result[T]) Error() error { return r.e }
func (r Result[T]) Value() T     { return r.v }

func (r Result[T]) MarshalJSON() ([]byte, error) {
    if r.e != nil {
        return json.Marshal(map[string]interface{}{"err": r.e.Error()})
    }
    return json.Marshal(map[string]interface{}{"val": r.v})
}
