package domain

import "fmt"

type Task struct {
	ID int
	Title string
	Description string
	Done bool
}

var (
	ErrTitleEmpty = fmt.Errorf("title is empty")
	ErrDescriptionEmpty = fmt.Errorf("description is empty")
)

func CreateTask(
	id int,
	title string,
	description string,
) (*Task, error) {
	// Провалидируем данные
	if title == "" {
		return nil, ErrTitleEmpty
	}
	if description == "" {
		return nil, ErrDescriptionEmpty
	}

	return &Task{
		ID: id,
		Title: title,
		Description: description,
		Done: false,
	}, nil
}