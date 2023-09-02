package auth

import (
	"log"

	"github.com/gh0stl1m/subscription-service/drivers/shared"
	"github.com/go-chi/chi/v5"
  "github.com/gh0stl1m/subscription-service/domains/users"
  "github.com/gh0stl1m/subscription-service/domains/auth"
)

type AuthCtx struct {
  UserServices users.IUserService
  AuthServices auth.IAuthServices
  InfoLog *log.Logger
  ErrorLog *log.Logger
}

func NewAuthRouter(app AuthCtx) shared.RouteConfigurer {

  return &AuthCtx {
    AuthServices: app.AuthServices,
    UserServices: app.UserServices,
    InfoLog: app.InfoLog,
    ErrorLog: app.ErrorLog,
  }
}

func (ac *AuthCtx) ConfigureRouter(r chi.Router) {

  r.Route("/auth", func (acr chi.Router) {

    acr.Post("/register", ac.RegisterHandler)
    acr.Post("/login", ac.LoginHandler)
  })
}
