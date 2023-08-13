package postgres

import (
	userDomain "github.com/gh0stl1m/subscription-service/domains/users"
	"gorm.io/gorm"
)

type UserRepository struct {
  db *gorm.DB
}

func NewUserRespository(db *gorm.DB) userDomain.UserRepository {

  return &UserRepository { db }
}

func (ur *UserRepository) Insert(user userDomain.User) error  {

  result := ur.db.Create(&user)

  return result.Error
}

func (ur *UserRepository) Find() ([]*userDomain.User, error) {

  users := []*userDomain.User{}

  result := ur.db.Find(&users)

  if result.Error != nil {

    return nil, result.Error
  }

  return users, nil
}

func (ur *UserRepository) FindOneBy(conditions userDomain.User) (*userDomain.User, error) {

  user := userDomain.User{}

  result := ur.db.Model(conditions).First(&user)

  if result.Error != nil {

    return nil, result.Error
  }

  return &user, nil
}

func (ur *UserRepository) DeleteById(id string) error {
  
  user := userDomain.User{}

  result := ur.db.Delete(&user, id)

  if result.Error != nil {

    return result.Error
  }

  return nil
}

func (ur *UserRepository) Update(columnsToChange userDomain.User) error {

  user := userDomain.User{}

  result := ur.db.Model(&user).Updates(columnsToChange)

  if result.Error != nil {

    return result.Error
  }

  return nil
}
