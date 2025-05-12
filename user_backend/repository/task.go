package repository

import (
	"user_backend/domain"

	"github.com/google/uuid"
)

type TaskDB interface {
	GetStatusByID(uuid uuid.UUID) (string, error)
	SetTask(task *domain.Task) error
	GetTaskByUUID(uuid uuid.UUID) (*domain.Task, error)
}
