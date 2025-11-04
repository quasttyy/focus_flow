package domain

import "context"

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(ctx context.Context, title, description string) (*Task, error) {
	task, err := CreateTask(0, title, description)
	if err != nil {
		// TODO: log
		return nil, err
	}
	id, err := s.repo.Create(ctx, task)
}