package tasks

import (
	"time"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)


type Task struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title" validate:"required,max=100"`
	Description string    `json:"description" validate:"required"`
	Status      string    `json:"status" validate:"required,oneof=todo inprogress done"`
	Priority    string    `json:"priority" validate:"required,oneof=low medium high"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Assignee    string    `json:"assignee" validate:"required"`
	AuthorID   uuid.UUID   `json:"author_id"`
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
