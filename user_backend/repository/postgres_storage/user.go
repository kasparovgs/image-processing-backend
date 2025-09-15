package postgres_storage

import (
	"database/sql"
	"user_backend/domain"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(connStr string) (*UserDB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &UserDB{db: db}, nil
}

func (ps *UserDB) RegisterUser(userID string, login string, psswrd string) error {
	query := "INSERT INTO users (id, login, password) VALUES ($1, $2, $3)"
	_, err := ps.db.Exec(query, userID, login, psswrd)
	if err != nil {
		return err
	}
	return nil
}

func (ps *UserDB) GetUserIDByLogin(login string) (string, error) {
	var userID string
	query := "SELECT id FROM users WHERE login = $1"
	err := ps.db.QueryRow(query, login).Scan(&userID)
	if err == sql.ErrNoRows {
		return "", domain.ErrNotFound("user not found")
	}
	return userID, nil
}

func (ps *UserDB) GetUserByUserID(userID string) (*domain.User, error) {
	var u domain.User
	query := "SELECT id, login, password FROM users WHERE id = $1"
	err := ps.db.QueryRow(query, userID).Scan(&u.UserID, &u.Login, &u.Password)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (ps *UserDB) IsAlreadyExist(login string) bool {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)"
	_ = ps.db.QueryRow(query, login).Scan(&exist)
	return exist
}
