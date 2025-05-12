package ram_storage

import (
	"user_backend/domain"

	"github.com/google/uuid"
)

type TaskDB struct {
	taskByUUID    map[uuid.UUID]*domain.Task // uuid -> domain.Task
	taskByOwnerID map[string]*domain.Task    // sessionID -> domain.Task
}

func NewTaskDB() *TaskDB {
	return &TaskDB{
		taskByUUID:    make(map[uuid.UUID]*domain.Task),
		taskByOwnerID: make(map[string]*domain.Task),
	}
}

func (rs *TaskDB) GetStatusByID(key uuid.UUID) (string, error) {
	task, exist := rs.taskByUUID[key]
	if !exist {
		return "", domain.ErrNotFound("task not found")
	}
	return task.Status, nil
}

func (rs *TaskDB) SetTask(task *domain.Task) error {
	if _, exist := rs.taskByUUID[task.UUID]; exist {
		return domain.ErrAlreadyExist("task already exist")
	}
	rs.taskByUUID[task.UUID] = task
	rs.taskByOwnerID[task.UserID] = task
	return nil
}

func (rs *TaskDB) GetTaskByUUID(uuid uuid.UUID) (*domain.Task, error) {
	task, exist := rs.taskByUUID[uuid]
	if !exist {
		return nil, domain.ErrNotFound("task not found")
	}
	return task, nil
}
