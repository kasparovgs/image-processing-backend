package usecases

type User interface {
	RegisterUser(username string, password string) (string, error)
	LoginUser(login string, password string) (string, error)
}
