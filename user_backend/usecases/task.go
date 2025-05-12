package usecases

import (
	"user_backend/domain"

	"github.com/google/uuid"
)

type Task interface {
	GetStatus(uuid uuid.UUID, sessionID string) (string, error)
	GetResult(uuid uuid.UUID, sessionID string) (string, error)
	PostTask(base64Image string, filter domain.Filter, sessionID string) (uuid.UUID, error)
	CommitTask(task *domain.Task) error
}
