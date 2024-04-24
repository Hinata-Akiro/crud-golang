package users

import (
	"time"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `db:"id" json:"id"`
	FirstName  string    `db:"first_name" json:"first_name" validate:"required"`
	LastName   string    `db:"last_name" json:"last_name" validate:"required"`
	Email      string    `db:"email" json:"email" validate:"required,email"`
	PassWord   string    `db:"password" json:"password" validate:"required"`
    CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type LoginInput struct {
    Email string
    Password string
}

func (u *LoginInput) Validate () error {
    validate := validator.New()
    return validate.Struct(u)
}



func (u *User) Validate () error {
	validate := validator.New()
	return validate.Struct(u)
}