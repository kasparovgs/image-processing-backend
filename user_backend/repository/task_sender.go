package repository

import (
	"user_backend/domain"
)

type TaskSender interface {
	Send(task domain.Task) error
}
