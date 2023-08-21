package users

import (
	"gorm.io/gorm"
)

type UserRepository struct {
  db *gorm.DB
}

func NewUserRespository(db *gorm.DB) IUserRepository {

  return &UserRepository { db }
}

func (ur *UserRepository) Insert(user User) error  {

  result := ur.db.Create(&user)

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

func (ur *UserRepository) DeleteById(id string) error {
  
  user := User{}

  result := ur.db.Delete(&user, id)

  if result.Error != nil {

    return result.Error
  }

  return nil
}

func (ur *UserRepository) Update(columnsToChange User) error {

  user := User{}

  result := ur.db.Model(&user).Updates(columnsToChange)

  if result.Error != nil {

    return result.Error
  }

  return nil
}
