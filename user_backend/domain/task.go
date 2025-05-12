package domain

import "github.com/google/uuid"

type Task struct {
	UserID      string    `json:"owner"`
	UUID        uuid.UUID `json:"uuid"`
	Status      string    `json:"status"`
	Base64Image string    `json:"base64image"`
	Filter      Filter    `json:"filter"`
}

type Filter struct {
	Name       string         `json:"name"`
	Parameters map[string]any `json:"parameters"`
}

const InProcess = "in_process"
const Ready = "ready"
