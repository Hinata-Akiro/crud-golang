package tasks

import (
	"time"
	"github.com/go-playground/validator/v10"
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

func dateValidation(fl validator.FieldLevel) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

func (a *Task) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("due_date", dateValidation)
	return validate.Struct(a)
}
