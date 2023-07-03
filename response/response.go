package response

import (
    "encoding/json"
    "errors"
    
    "github.com/go4s/util/result"
)

type (
    inner[T any] struct {
        Code int    `json:"code"`
        Msg  string `json:"msg"`
        Data T      `json:"data"`
    }
    Response[T any] struct {
        c int
        m string
        v T
    }
)

func (o Response[T]) Code() int       { return o.c }
func (o Response[T]) Message() string { return o.m }
func (o Response[T]) Data(matches ...int) result.Result[T] {
    success := true
    for i := range matches {
        if success = o.c == matches[i]; success {
            break
        }
    }
    if success {
        return result.Success[T](o.v)
    }
    return result.Failed[T](errors.New(o.m))
}
func (o Response[T]) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{"code": o.c, "msg": o.m, "data": o.v})
}

func (o *Response[T]) UnmarshalJSON(raw []byte) (err error) {
    var resp inner[T]
    if err = json.Unmarshal(raw, &resp); err != nil {
        return
    }
    o.c, o.m, o.v = resp.Code, resp.Msg, resp.Data
    return
    
}
func (o *Response[T]) Failed(code int, e error) *Response[T] { o.c, o.m = code, e.Error(); return o }
func (o *Response[T]) Success(code int, v T) *Response[T]    { o.c, o.v = code, v; return o }

func New[T any]() *Response[T]                       { return &Response[T]{} }
func Failed[T any](code int, err error) *Response[T] { return New[T]().Failed(code, err) }
func Success[T any](code int, val T) *Response[T]    { return New[T]().Success(code, val) }
