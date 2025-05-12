package usecases

type Session interface {
	CreateSessionID() (string, error)
}
