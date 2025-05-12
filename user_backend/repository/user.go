package repository

import "user_backend/domain"

type User interface {
	IsAlreadyExist(userID string) bool
	RegisterUser(userID string, login string, psswrd string) error
	GetUserIDByLogin(login string) (string, error)
	GetUserByUserID(userID string) (*domain.User, error)
}
