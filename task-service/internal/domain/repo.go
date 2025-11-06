package domain

import (
	"context"
)

type TaskRepository interface {
	InsertTask(ctx context.Context, task *Task) (int, error)
	ListTasks(ctx context.Context) ([]*Task, error)
	CompleteTask(ctx context.Context, id int) error
}