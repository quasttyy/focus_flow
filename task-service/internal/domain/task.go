package domain

type Task struct {
	ID int
	Title string
	Description string
	Done bool
}

func CreateTask()