package repo

import (
	"context"
	"focus_flow/common/logger"
	"focus_flow/task-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	Pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{Pool: pool}
}

const InsertTaskQuery = `
INSERT INTO tasks (title, description, done)
VALUES ($1, $2, $3)
RETURNING id
`

func (r *PostgresRepo) InsertTask(ctx context.Context, task *domain.Task) (int, error) {
	var id int
	err := r.Pool.QueryRow(ctx, InsertTaskQuery, task.Title, task.Description, task.Done).Scan(&id)
	if err != nil {
		logger.Log.Error("failed to insert task:", task.Title)
		return 0, err
	}

	return id, nil
}

const ListTasksQuery = `
SELECT id, title, description, done
FROM tasks
`

func (r *PostgresRepo) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	var tasks []*domain.Task
	rows, err := r.Pool.Query(ctx, ListTasksQuery)
	if err != nil {
		logger.Log.Error("failed to query tasks:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
			logger.Log.Error("error scanning row", err)
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Error("rows iteration error:", err)
		return nil, err
	}

	logger.Log.Info("retrieved", len(tasks), "tasks from DB")

	return tasks, nil 
}