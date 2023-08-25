package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gh0stl1m/subscription-service/domains/users"
	"github.com/gh0stl1m/subscription-service/utils"
)

type RegisterResponse struct {
	Message string `json:"message,omitempty"`
}

func (ac *AuthCtx) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userBody users.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userBody)

	if err != nil {

		ac.ErrorLog.Println("User validation failed")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(RegisterResponse{Message: "Invalid user"})

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
		json.NewEncoder(w).Encode(RegisterResponse{Message: err.Error()})

		return
	}

	if err != nil {

		ac.ErrorLog.Println("Something went wrong creating user")

		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(RegisterResponse{Message: err.Error()})

		return
	}

	w.WriteHeader(http.StatusCreated)
}
