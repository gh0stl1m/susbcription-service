package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServices struct {}

func NewAuthServices() IAuthServices {

  return &AuthServices{}
}

func (as *AuthServices) GenerateToken(username string) (string, error) {

  secret := []byte(os.Getenv("TOKEN_SECRET"))
  token := jwt.New(jwt.SigningMethodHS256)

  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(10 * time.Minute)
  claims["authorized"] = true
  claims["username"] = username

  tokenString, err := token.SignedString(secret)

  if err != nil {

    return "", nil
  }

  return tokenString, nil
}
