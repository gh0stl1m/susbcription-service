package users

import "time"

type User struct {
  ID string `gorm:"primaryKey"`
  Email string
  FirstName string
  LastName string
  Password string
  Active uint
  IsAdmin uint
  CreatedAt time.Time
  UpdatedAt time.Time
}

type UserRepository interface {
  Insert(user User) (string, error)
  FindOne(id string) (*User, error)
  Find() ([]*User, error)
  FindByEmail(email string) (*User, error)
  DeleteById(id string) error
  ResetPassword(password string) error
}
