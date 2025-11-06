package domain

import "fmt"

type Task struct {
	ID int
	Title string
	Description string
	Done bool
}

var (
	ErrEmptyTitle = fmt.Errorf("title is empty")
)

func CreateTask(
	title string,
	description string,
) (*Task, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	return &Task{
		Title: title,
		Description: description,
		Done: false,
	}, nil
}