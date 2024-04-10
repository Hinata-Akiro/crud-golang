package tasks

import (
	"errors"
	"time"
)

type Task struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date"`
	Assignee    string    `json:"assignee"`
}

func ValidateTaskInputs(title, description, status, priority string, dueDate time.Time, assignee string) error {

	if title == "" {
		return errors.New("title is required")
	}

	if description == "" {
		return errors.New("description is required")
	}

	if status == "" {
		return errors.New("status is required")
	}

	if priority == "" {
		return errors.New("priority is required")
	}

	if dueDate.IsZero() {
		return errors.New("due date is required")
	}

	if assignee == "" {
		return errors.New("assignee is required")
	}

	return nil
}

func NewTask(title, description, status, priority string, dueDate time.Time, assignee string) (*Task, error) {
	err := ValidateTaskInputs(title, description, status, priority, dueDate, assignee)

	if err != nil {
		panic(err)
	}

	now := time.Now().UTC()

	task := &Task{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		CreatedAt:   now,
		UpdatedAt:   now,
		DueDate:     dueDate,
		Assignee:    assignee,
	}

	return task, nil
}
