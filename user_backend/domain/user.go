package domain

type User struct {
	UserID   string `json:"userID"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
