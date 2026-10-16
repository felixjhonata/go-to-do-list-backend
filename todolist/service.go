package todolist

type todoitem struct {
	ID int `json:"id"`
	Task string `json:"task"`
	Assignee string `json:"assignee"`
}

var todoitems = []todoitem {
	{ID: 0, Task: "Vacuum the floor", Assignee: "Stan"},
	{ID: 1, Task: "Wipe the floor", Assignee: "Stacy"},
	{ID: 2, Task: "Cook Dinner", Assignee: "Mom"},
	{ID: 3, Task: "Buy Grocery", Assignee: "Dad"},
}

func getToDoItems() []todoitem {
	return todoitems
}