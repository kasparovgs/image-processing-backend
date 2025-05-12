package http

import (
	"net/http"
	"user_backend/api/http/types"
	"user_backend/usecases"

	"github.com/go-chi/chi/v5"
)

// Task represents an HTTP handler for managing tasks.
type Task struct {
	service usecases.Task
}

// NewHandler creates a new instance of Task.
func NewTaskHandler(service usecases.Task) *Task {
	return &Task{service: service}
}

// @Summary Get status of a task
// @Security SessionIDAuth
// @Description Get status of a task by its uuid
// @Tags task
// @Accept  json
// @Produce json
// @Param task_id path string true "UUID of the task"
// @Success 200 {string} types.GetTaskStatusHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Failure 401 {string} string "Unauthorized"
// @Router /status/{task_id} [get]
func (t *Task) getStatusHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetTaskStatusHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	status, err := t.service.GetStatus(req.UUID, req.SessionID)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	types.ProcessError(w, err, &types.GetTaskStatusHandlerResponse{Status: status})
}

// @Summary Submit the task for processing and returns task_id
// @Security SessionIDAuth
// @Description Submit the task with image upload and returns task_id (uint64)
// @Tags task
// @Accept json
// @Produce json
// @Param image body types.PostTaskHandlerRequest true "Image in base64 and filters"
// @Success 201 {object} types.PostTaskHandlerResponse "The task is running successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /task [post]
func (t *Task) postTaskHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostTaskHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	uuid, err := t.service.PostTask(req.Base64Image, req.Filter, req.SessionID)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	w.WriteHeader(http.StatusCreated)
	types.ProcessError(w, err, &types.PostTaskHandlerResponse{Task_id: uuid})
}

// @Summary Get result of a task
// @Security SessionIDAuth
// @Description Get result of a task by its uuid
// @Tags task
// @Accept  json
// @Produce json
// @Param task_id path string true "UUID of the task"
// @Success 200 {object} types.GetTaskResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 400 {string} string "the task is still in process"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Task not found"
// @Router /result/{task_id} [get]
func (t *Task) getResultHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetTaskResultHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	result, err := t.service.GetResult(req.UUID, req.SessionID)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	types.ProcessError(w, err, &types.GetTaskResultHandlerResponse{Base64Image: result})
}

func (t *Task) commitTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := types.CreateCommitTaskHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
	err = t.service.CommitTask(task)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}
}

func (t *Task) WithTaskHandlers(r chi.Router) {
	r.Post("/task", t.postTaskHandler)
	r.Get("/status/{task_id}", t.getStatusHandler)
	r.Get("/result/{task_id}", t.getResultHandler)
	r.Post("/commit", t.commitTaskHandler)
}
