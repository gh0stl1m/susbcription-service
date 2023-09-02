package users

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

  "github.com/gh0stl1m/subscription-service/utils"
)

type UserServices struct {
	Repository IUserRepository
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func NewUserService(repository IUserRepository, infoLog, errorLog *log.Logger) IUserService {

	return &UserServices{Repository: repository, InfoLog: infoLog, ErrorLog: errorLog}
}

func (us *UserServices) FindOneBy(conditions User, selector []string) (*User, error) {

  return us.Repository.FindOneBy(conditions, selector)
}

func (us *UserServices) Create(user UserDTO) error {

	userData, err := us.Repository.FindOneBy(User{Email: user.Email}, []string{"id"})

	if err != nil {

		us.ErrorLog.Println("Error checking user")

		return err
	}

	if userData != nil {

    return fmt.Errorf("%w", utils.ErrUserAlreadyExists)
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {

		us.ErrorLog.Println("Error hashing password")

		return err
	}

	user.Password = string(passwordHashed)

	return us.Repository.Insert(user)
}

func (us *UserServices) PasswordMatches(hash, plainText string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))

	return err == nil
}

func (us *UserServices) ResetPassword(id uuid.UUID, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {

		us.ErrorLog.Println("Error hashing passoword")

		return err
	}
	condition := User{Password: string(hashedPassword)}

	result := us.Repository.Update(id, condition)

	if result != nil {

		us.ErrorLog.Println("Error updating user")

		return result
	}

	return nil
}
