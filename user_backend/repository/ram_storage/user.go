package ram_storage

import (
	"user_backend/domain"
)

type UserDB struct {
	usersByID      map[string]domain.User // userID -> User
	usersIDByLogin map[string]string      // login -> userID
}

func NewUserDB() *UserDB {
	return &UserDB{
		usersByID:      make(map[string]domain.User),
		usersIDByLogin: make(map[string]string),
	}
}

func (rs *UserDB) RegisterUser(userID string, login string, psswrd string) error {
	rs.usersByID[userID] = domain.User{
		UserID:   userID,
		Login:    login,
		Password: psswrd,
	}

	rs.usersIDByLogin[login] = userID
	return nil
}

func (rs *UserDB) GetUserIDByLogin(login string) (string, error) {
	userID, exist := rs.usersIDByLogin[login]
	if !exist {
		return "", domain.ErrNotFound("user not found")
	}
	return userID, nil
}

func (rs *UserDB) GetUserByUserID(userID string) (*domain.User, error) {
	user, exist := rs.usersByID[userID]
	if !exist {
		return nil, domain.ErrNotFound("user not found")
	}
	return &user, nil
}

func (rs *UserDB) IsAlreadyExist(login string) bool {
	_, exist := rs.usersIDByLogin[login]
	return exist
}
