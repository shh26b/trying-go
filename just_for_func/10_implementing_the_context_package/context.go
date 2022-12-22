package contextimpl

import "time"

type Context interface {
	DeadLine() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
