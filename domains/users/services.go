package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	Repository IUserRepository
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func NewUserService(repository IUserRepository) IUserService {

	return &UserServices{repository}
}

func (us *UserServices) FindOneByEmail(email string) (*User, error) {

	user, err := us.Repository.FindOneBy(User{Email: email})

	if err != nil {

		us.ErrorLog.Println("Something went wrong reading user")

		return nil, err
	}


	return user, err
}

func (us *UserServices) PasswordMatches(hash, plainText string) bool {

  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))

  if err != nil {

    return false
  }

  return true
}

func (us *UserServices) ResetPassword(id, password string) error {

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  
	if err != nil {

    us.ErrorLog.Println("Error hashing passoword")
    
		return err
	}
  condition := User{ Password: string(hashedPassword) }

  result := us.Repository.Update(id, condition)

  if result != nil {
    
    us.ErrorLog.Println("Error updating user")

    return result
  }

  return nil
}

