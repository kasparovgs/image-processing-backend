package postgres_storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"user_backend/domain"

	"github.com/google/uuid"
)

type TaskDB struct {
	db *sql.DB
}

func NewTaskDB(connStr string) (*TaskDB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &TaskDB{db: db}, nil
}

func (ps *TaskDB) GetStatusByID(key uuid.UUID) (string, error) {
	var status string
	err := ps.db.QueryRow("SELECT status FROM tasks WHERE key = $1", key).Scan(&status)
	if err == sql.ErrNoRows {
		return "", domain.ErrNotFound("task not found")
	}
	if err != nil {
		return "", err
	}

	return status, nil
}

func (ps *TaskDB) SetTask(task *domain.Task) error {
	params, err := json.Marshal(task.Filter.Parameters)
	if err != nil {
		return err
	}

	query := `INSERT INTO tasks (uuid, user_id, status, base64image, filter_name, filter_parameters) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = ps.db.Exec(query, task.UUID,
		task.UserID, task.Status,
		task.Base64Image, task.Filter.Name, params)
	return err
}

func (ps *TaskDB) GetTaskByUUID(uuid uuid.UUID) (*domain.Task, error) {
	task := &domain.Task{}
	var paramsJSON []byte
	query := `SELECT uuid, user_id, status, base64image, filter_name, filter_parameters
			  FROM tasks WHERE uuid = $1`

	err := ps.db.QueryRow(query, uuid).Scan(&task.UUID,
		&task.UserID,
		&task.Status,
		&task.Base64Image,
		&task.Filter.Name,
		&paramsJSON)

	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound("task not found")
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(paramsJSON, &task.Filter.Parameters); err != nil {
		return nil, fmt.Errorf("unmarshal filter parameters: %w", err)
	}
	return task, nil
}

func (ps *TaskDB) UpdateTask(task *domain.Task) error {
	if task == nil {
		return fmt.Errorf("invalid parameters")
	}

	params, err := json.Marshal(task.Filter.Parameters)
	if err != nil {
		return err
	}

	query := `UPDATE tasks SET user_id = $2, status = $3, base64image = $4,
			  filter_name = $5, filter_parameters = $6
        	  WHERE uuid = $1`
	_, err = ps.db.Exec(query,
		task.UUID, task.UserID, task.Status, task.Base64Image, task.Filter.Name, params,
	)
	return err
}
