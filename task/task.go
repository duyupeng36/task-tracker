package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ErrNilTask = errors.New("task can't be nil")

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format("2006-01-02 15:04:05 -0700"))
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	var timeString string
	err = json.Unmarshal(data, &timeString)
	if err != nil {
		return err
	}

	t.Time, err = time.Parse("2006-01-02 15:04:05 -0700", timeString)
	return err
}

type Task struct {
	ID          int       `json:"id"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreateAt    JSONTime  `json:"create_at"`
	UpdateAt    *JSONTime `json:"update_at"`
}

func (task *Task) Update(status string, description string) error {
	if task == nil {
		return ErrNilTask
	}

	if status != "" {
		task.Status = status
	}

	if description != "" {
		task.Description = description
	}

	task.UpdateAt = &JSONTime{time.Now()}
	return nil
}

func (task *Task) String() string {
	if task == nil {
		return "<nil>"
	}

	var createAt, updateAt string

	createAt = task.CreateAt.Format("2006-01-02 15:04:05 -0700")

	if task.UpdateAt == nil {
		updateAt = "nil"
	} else {
		updateAt = task.UpdateAt.Format("2006-01-02 15:04:05 -0700")
	}

	return fmt.Sprintf("Task(id=%d, status=%s, description=%s, create_at=%v, update_at=%v)", task.ID, task.Status, task.Description, createAt, updateAt)
}

func (task *Task) GoString() string {
	return task.String()
}
