package service

import (
	"errors"
	"fmt"
	"user_backend/domain"
	"user_backend/repository"

	"github.com/google/uuid"
)

type Task struct {
	taskRepo    repository.TaskDB
	sessionRepo repository.Session
	sender      repository.TaskSender
}

func NewTask(taskRepo repository.TaskDB, sessionRepo repository.Session, sender repository.TaskSender) *Task {
	return &Task{
		taskRepo:    taskRepo,
		sessionRepo: sessionRepo,
		sender:      sender,
	}
}

func (t *Task) GetStatus(uuid uuid.UUID, sessionID string) (string, error) {
	task, err := t.taskRepo.GetTaskByUUID(uuid)
	if err != nil {
		return "", err
	}
	userId, err := t.sessionRepo.GetUserIDBySessionID(sessionID)
	if err != nil {
		return "", err
	}

	if task.UserID != userId {
		return "", domain.ErrForbidden("wrong userID")
	}
	return task.Status, nil
}

func (t *Task) GetResult(uuid uuid.UUID, sessionID string) (string, error) {
	task, err := t.taskRepo.GetTaskByUUID(uuid)
	if err != nil {
		return "", err
	}

	userId, err := t.sessionRepo.GetUserIDBySessionID(sessionID)
	if err != nil {
		return "", err
	}

	if task.UserID != userId {
		return "", domain.ErrForbidden("wrong userID")
	}

	if status := task.Status; status == domain.InProcess {
		return "", errors.New("the task is still in process")
	}
	base64img := task.Base64Image
	return base64img, nil
}

func (t *Task) PostTask(base64Image string, filter domain.Filter, sessionID string) (uuid.UUID, error) {
	userId, err := t.sessionRepo.GetUserIDBySessionID(sessionID)
	if err != nil {
		return uuid.Nil, err
	}

	id := uuid.New()
	task := domain.Task{UserID: userId, UUID: id, Status: domain.InProcess, Base64Image: base64Image, Filter: filter}
	err = t.sender.Send(task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("sending task: %w", err)
	}
	t.taskRepo.SetTask(&task)
	return id, nil
}

func (t *Task) CommitTask(task *domain.Task) error {
	oldTask, err := t.taskRepo.GetTaskByUUID(task.UUID)
	if err != nil {
		return err
	}
	oldTask.Base64Image = task.Base64Image
	oldTask.Status = task.Status
	return nil
}
