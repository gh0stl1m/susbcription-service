package http

import (
	"net/http"

	"github.com/gh0stl1m/subscription-service/drivers/http/healthcheks"
	"github.com/gh0stl1m/subscription-service/drivers/shared"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

  configurers := []shared.RouteConfigurer{
    healtcheckRouter,
  }

  for _, configurer := range configurers {
      configurer.ConfigureRouter(mux)
  }

  return mux
}
