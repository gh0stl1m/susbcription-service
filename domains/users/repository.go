package users

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRespository(db *gorm.DB) IUserRepository {

	return &UserRepository{db}
}

func (ur *UserRepository) Insert(user UserDTO) error {

	newUser := User{
		ID:         uuid.NewV4(),
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Password:   user.Password,
		UserActive: user.Active,
		IsAdmin:    user.IsAdmin,
	}

	result := ur.db.Create(&newUser)

	return result.Error
}

func (ur *UserRepository) FindOneBy(conditions User, selector []string) (*User, error) {

	user := User{}

	result := ur.db.Select(selector).Where(conditions).First(&user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return nil, nil
		}

		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) Update(id uuid.UUID, columnsToChange User) error {

	user := User{ID: id}

	result := ur.db.Model(&user).Updates(columnsToChange)

	if result.Error != nil {

		return result.Error
	}

	return nil
}
