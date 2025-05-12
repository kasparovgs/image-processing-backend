package service

import (
	"errors"
	"pkg/hash"
	"user_backend/domain"
	"user_backend/repository"

	"github.com/google/uuid"
)

type User struct {
	userRepo    repository.User
	sessionRepo repository.Session
}

func NewUser(userRepo repository.User, sessionRepo repository.Session) *User {
	return &User{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (u *User) RegisterUser(login string, password string) (string, error) {
	if u.userRepo.IsAlreadyExist(login) {
		return "", domain.ErrAlreadyExist("user with this login is already exist")
	}

	hashStr, err := hash.CreateHash(password)
	if err != nil {
		return "", errors.New("wrong type of password")
	}

	sessionID, err := CreateSessionID()
	if err != nil {
		return "", err
	}

	userID := uuid.New().String()
	u.userRepo.RegisterUser(userID, login, hashStr)

	u.sessionRepo.SetUserIDBySessionID(sessionID, userID)
	return sessionID, nil
}

func (u *User) LoginUser(login string, password string) (string, error) {
	if !u.userRepo.IsAlreadyExist(login) {
		return "", domain.ErrNotFound("user not found")
	}

	hashStr, err := hash.CreateHash(password)
	if err != nil {
		return "", errors.New("wrong type of password")
	}

	userID, err := u.userRepo.GetUserIDByLogin(login)
	if err != nil {
		return "", err
	}

	user, err := u.userRepo.GetUserByUserID(userID)
	if err != nil {
		return "", err
	}
	if login != user.Login || hashStr != user.Password {
		return "", domain.ErrUnauthorized("incorrect login or password")
	}

	sessionID, err := CreateSessionID()
	if err != nil {
		return "", err
	}

	u.sessionRepo.SetUserIDBySessionID(sessionID, userID)
	return sessionID, nil
}
