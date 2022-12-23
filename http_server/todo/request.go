package todo

import "errors"

type createTodo struct {
	Todo string
}

func (t *createTodo) validate() error {
	l := len(t.Todo)
	if l < 2 || l > 100 {
		return errors.New("todo must has to " +
			" be at least 2 characters" +
			" and under 100 characters")
	}
	return nil
}
