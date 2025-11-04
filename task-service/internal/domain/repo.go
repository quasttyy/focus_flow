package domain

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, ) (int, error)
	ListTasks(ctx context.Context, user_id string) ([]*Task, error)
	CompleteTask(ctx context.Context, id int) error
}
