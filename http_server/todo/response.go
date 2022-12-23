package todo

import "fmt"

type Todo struct {
	Id          string
	Content     string
	IsCompleted bool
}

func (t *Todo) String() string {
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
