package http

import (
	"net/http"
	"user_backend/api/http/types"
	"user_backend/usecases"

	"github.com/go-chi/chi/v5"
)

type User struct {
	service usecases.User
}

func NewUserHandler(service usecases.User) *User {
	return &User{service: service}
}

// @Summary Registers a new user
// @Description Registers a new user and issue their sessionID
// @Tags user
// @Accept  json
// @Produce json
// @Param request body types.PostRegisterUserHandlerRequest true "login and password"
// @Success 201 {string} types.PostRegisterUserHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 409 {string} string "User with this username already exist"
// @Router /register [post]
func (u *User) postRegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostRegisterUserHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	token, err := u.service.RegisterUser(req.Login, req.Password)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	w.WriteHeader(http.StatusCreated)
	types.ProcessError(w, err, &types.PostRegisterUserHandlerResponse{SessionID: token})
}

// @Summary Logs in a user
// @Description Authenticates a user with login and password, returns a sessionID
// @Tags user
// @Accept  json
// @Produce json
// @Param request body types.PostRegisterUserHandlerRequest true "Login and password"
// @Success 200 {object} types.PostRegisterUserHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Incorrect login or password"
// @Failure 404 {string} string "User not found"
// @Router /login [post]
func (u *User) postLoginUserHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostRegisterUserHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	token, err := u.service.LoginUser(req.Login, req.Password)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	types.ProcessError(w, err, &types.PostRegisterUserHandlerResponse{SessionID: token})
}

func (u *User) WithUserHandlers(r chi.Router) {
	r.Post("/register", u.postRegisterUserHandler)
	r.Post("/login", u.postLoginUserHandler)
}
