package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  
  "github.com/gh0stl1m/subscription-service/drivers/http/healthcheks"
	"github.com/gh0stl1m/subscription-service/drivers/http/auth"
	"github.com/gh0stl1m/subscription-service/domains/users"
	"github.com/gh0stl1m/subscription-service/drivers/shared"

)

func (app *Config) Router() http.Handler {

  mux := chi.NewRouter()

  mux.Use(middleware.Recoverer)
  mux.Use(app.LoadSession)

  healtcheckRouter := healthcheks.NewHealthcheckRouter(healthcheks.HealthCheckCtx {
      DB: app.DB,
      ErrorLog: app.ErrorLog,
      InfoLog: app.InfoLog,
    },
  )

  userRepository := users.NewUserRespository(app.DB)
  userService := users.NewUserService(userRepository, app.InfoLog, app.ErrorLog)

  authRouter := auth.NewAuthRouter(auth.AuthCtx{
    UserServices: userService,
    InfoLog: app.InfoLog,
    ErrorLog: app.ErrorLog,
  })

  configurers := []shared.RouteConfigurer{
    healtcheckRouter,
    authRouter,
  }

  for _, configurer := range configurers {
      configurer.ConfigureRouter(mux)
  }

  return mux
}
