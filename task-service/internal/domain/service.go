package domain

import (
	"context"
	"focus_flow/common/logger"
)

type TaskService struct {
	Repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (ts *TaskService) Create(
	ctx context.Context,
	title,
	description string,
) (*Task, error) {
	task, err := CreateTask(title, description)
	if err != nil {
		logger.Log.Error("error creating task:", task.Title)
		return nil, err
	}

	id, err := ts.Repo.InsertTask(ctx, task)
	if err != nil {
		logger.Log.Error("error while insert task:", err)
		return nil, err
	}

	task.ID = id
	logger.Log.Info("task:", task.Title, "with id:", task.ID, "successfully created")

	return task, nil
}

func (ts *TaskService) List(ctx context.Context) ([]*Task, error) {
	tasks, err := ts.Repo.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}