package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gh0stl1m/subscription-service/domains/auth"
	"github.com/gh0stl1m/subscription-service/domains/users"
	"github.com/gh0stl1m/subscription-service/utils"
)

type ServiceResponse struct {
	Message string `json:"message,omitempty"`
}

type LoginResponse struct {
	AuthCode string `json:"authcode,omitempty"`
}

func (ac *AuthCtx) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userBody users.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userBody)

	if err != nil {

		ac.ErrorLog.Println("User validation failed")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ServiceResponse{Message: "Invalid user"})

		return
	}

	newUser := users.UserDTO{
		Email:     userBody.Email,
		FirstName: userBody.FirstName,
		LastName:  userBody.LastName,
		Password:  userBody.Password,
		Active:    1,
		IsAdmin:   userBody.IsAdmin,
	}

	err = ac.UserServices.Create(newUser)

	if errors.Is(err, utils.ErrUserAlreadyExists) {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ServiceResponse{Message: err.Error()})

		return
	}

	if err != nil {

		ac.ErrorLog.Println("Something went wrong creating user")

		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(ServiceResponse{Message: err.Error()})

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ac *AuthCtx) LoginHandler(w http.ResponseWriter, r *http.Request) {

  var loginBody auth.LoginRequestBody

  err := json.NewDecoder(r.Body).Decode(&loginBody)

  if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ServiceResponse{Message: "Invalid Payload"})

		return
  }

  user, err := ac.UserServices.FindOneBy(users.User{ Email: loginBody.Username }, []string{"id", "password"})

  if err != nil || user == nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ServiceResponse{Message: "Invalid Username or Password"})

		return
  }

  isPasswordMatches := ac.UserServices.PasswordMatches(user.Password, loginBody.Password)

  if !isPasswordMatches {
  
  	w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ServiceResponse{Message: "Invalid Username or Password"})

		return
  }

  authcode, err := ac.AuthServices.GenerateToken(loginBody.Username)  

  if err != nil || authcode == "" {

   	w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ServiceResponse{Message: "Error generating auth token"})

		return
  }
  
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(LoginResponse{AuthCode: authcode})
}
