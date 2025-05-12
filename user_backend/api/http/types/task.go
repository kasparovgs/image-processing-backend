package types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"user_backend/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type GetTaskStatusHandlerRequest struct {
	UUID      uuid.UUID `json:"uuid"`
	SessionID string    `json:"-"`
}

func CreateGetTaskStatusHandlerRequest(r *http.Request) (*GetTaskStatusHandlerRequest, error) {
	header := r.Header.Get("Authorization")
	sessionID, err := validateAccess(header)
	if err != nil {
		return nil, err
	}
	uuidStr := chi.URLParam(r, "task_id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	req := GetTaskStatusHandlerRequest{UUID: id, SessionID: sessionID}
	return &req, nil
}

type GetTaskStatusHandlerResponse struct {
	Status string `json:"status"`
}

type GetTaskResultHandlerRequest struct {
	UUID      uuid.UUID `json:"uuid"`
	SessionID string    `json:"-"`
}

func CreateGetTaskResultHandlerRequest(r *http.Request) (*GetTaskResultHandlerRequest, error) {
	header := r.Header.Get("Authorization")
	sessionID, err := validateAccess(header)
	if err != nil {
		return nil, err
	}
	uuidStr := chi.URLParam(r, "task_id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	req := GetTaskResultHandlerRequest{UUID: id, SessionID: sessionID}
	return &req, nil
}

type GetTaskResultHandlerResponse struct {
	Base64Image string `json:"result"`
}

type PostTaskHandlerRequest struct {
	Base64Image string        `json:"image"`
	Filter      domain.Filter `json:"filter"`
	SessionID   string        `json:"-"`
}

func CreatePostTaskHandlerRequest(r *http.Request) (*PostTaskHandlerRequest, error) {
	header := r.Header.Get("Authorization")
	sessionID, err := validateAccess(header)
	if err != nil {
		return nil, err
	}

	var req PostTaskHandlerRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	req.SessionID = sessionID
	return &req, nil
}

type PostTaskHandlerResponse struct {
	Task_id uuid.UUID `json:"task_id"`
}

func CreateCommitTaskHandlerRequest(r *http.Request) (*domain.Task, error) {
	var req domain.Task
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

func validateAccess(token string) (string, error) {
	if token == "" {
		return "", domain.ErrUnauthorized("missing authorization key")
	}

	if !strings.HasPrefix(token, "Bearer ") {
		return "", domain.ErrUnauthorized("invalid authorization format")
	}
	sessionID := strings.TrimPrefix(token, "Bearer ")
	return sessionID, nil
}
