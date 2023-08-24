package users

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
  db *gorm.DB
}

func NewUserRespository(db *gorm.DB) IUserRepository {

  return &UserRepository { db }
}

func (ur *UserRepository) Insert(user UserDTO) error  {

  newUser := User{
    ID: uuid.NewV4(),
    Email: user.Email,
    FirstName: user.FirstName,
    LastName: user.LastName,
    Password: user.Password,
    UserActive: user.Active,
    IsAdmin: user.IsAdmin,
  }

  result := ur.db.Create(&newUser)

  return result.Error
}

func (ur *UserRepository) Find() ([]*User, error) {

  users := []*User{}

  result := ur.db.Find(&users)

  if result.Error != nil {

    return nil, result.Error
  }

  return users, nil
}

func (ur *UserRepository) FindOneBy(conditions User) (*User, error) {

  user := User{}

  result := ur.db.Model(conditions).First(&user)

  if result.Error != nil {

    return nil, result.Error
  }

  return &user, nil
}

func (ur *UserRepository) DeleteById(id uuid.UUID) error {
  
  user := User{}

  result := ur.db.Delete(&user, id)

  if result.Error != nil {

    return result.Error
  }

  return nil
}

func (ur *UserRepository) Update(id uuid.UUID, columnsToChange User) error {

  user := User{ ID: id }

  result := ur.db.Model(&user).Updates(columnsToChange)

  if result.Error != nil {

    return result.Error
  }

  return nil
}
