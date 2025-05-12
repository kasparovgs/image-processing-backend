package usecases

import "user_backend/domain"

type Task interface {
	ProcessTask(task *domain.Task) error
	CommitTask(task *domain.Task) error
}
