package repository

type Session interface {
	GetUserIDBySessionID(sessionID string) (string, error)
	SetUserIDBySessionID(sessionID string, userID string) error
}
