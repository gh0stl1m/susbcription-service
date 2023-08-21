package auth

import (
	"github.com/gh0stl1m/subscription-service/drivers/shared"
	"github.com/go-chi/chi/v5"
)

type AuthCtx struct {}

func NewAuthRouter(app AuthCtx) shared.RouteConfigurer {

  return &AuthCtx {}
}

func (ac *AuthCtx) ConfigureRouter(r chi.Router) {

  r.Route("/authentication", func (acr chi.Router) {

    // acr.Post("/login")
    // acr.Post("/logout")
    // acr.Post("/register")
  })
}
