package main

import (
	"focus_flow/common/logger"
	"focus_flow/task-service/internal/postgres"
)

func main() {
	// Инициализируем логгер:
	
	// Подключимся к PostgreSQL
	dbPool, err := postgres.ConnectPostgres()
	if err != nil {
		logger.Log.Fatal("error connecting postgresql", err)
	}
	defer dbPool.Close()
}