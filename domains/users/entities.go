package users

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
  ID uuid.UUID `gorm:"type:uuid;primary_key;"`
  Email string 
  FirstName string
  LastName string
  Password string
  UserActive uint
  IsAdmin uint
  CreatedAt time.Time
  UpdatedAt time.Time
}

type UserDTO struct {
  Email string `json:"email"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  Active uint `json:"active,omitempty"`
  Password string `json:"password"`
  IsAdmin uint `json:"isAdmin"`
}

type IUserRepository interface {
  Insert(user UserDTO) error
  FindOneBy(conditions User, selector []string) (*User, error)
  Update(id uuid.UUID, columnsToChange User) error
}

type IUserService interface {
  Create(user UserDTO) error
  ResetPassword(id uuid.UUID, password string) error
  PasswordMatches(hash, plainText string) bool
}
