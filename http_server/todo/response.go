package todo

import "fmt"

type todo struct {
	Id          string
	Content     string
	IsCompleted bool
}

func (t *todo) String() string {
	return fmt.Sprintf(
		`{
			"id": "%v",
			"content": "%v",
			"is_completed": %v
		}`,
		t.Id,
		t.Content,
		t.IsCompleted,
	)
}
