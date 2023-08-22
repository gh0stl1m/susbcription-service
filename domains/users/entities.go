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

type IUserRepository interface {
  Insert(user User) error
  Find() ([]*User, error)
  FindOneBy(conditions User) (*User, error)
  DeleteById(id string) error
  Update(id string, columnsToChange User) error
}

type IUserService interface {
  ResetPassword(id, password string) error
  PasswordMatches(hash, plainText string) bool
  FindOneByEmail(email string) (*User, error)
}
