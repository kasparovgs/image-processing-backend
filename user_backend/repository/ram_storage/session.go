package ram_storage

import "user_backend/domain"

type Session struct {
	data map[string]string // sessionID -> userID
}

func NewSession() *Session {
	return &Session{
		data: make(map[string]string),
	}
}

func (rs *Session) GetUserIDBySessionID(sessionID string) (string, error) {
	userID, exist := rs.data[sessionID]
	if !exist {
		return "", domain.ErrNotFound("user not found")
	}
	return userID, nil
}

func (rs *Session) SetUserIDBySessionID(sessionID string, userID string) error {
	rs.data[sessionID] = userID
	return nil
}
