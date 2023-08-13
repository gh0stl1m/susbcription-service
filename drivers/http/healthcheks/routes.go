package healthcheks

import (
	"database/sql"
	"log"

	"github.com/gh0stl1m/subscription-service/drivers/shared"
	"github.com/go-chi/chi/v5"
)

type HealthCheckCtx struct {
  DB *sql.DB
  InfoLog *log.Logger
  ErrorLog *log.Logger
}

func NewHealthcheckRouter(app HealthCheckCtx) shared.RouteConfigurer {

  return &HealthCheckCtx{
    DB: app.DB,
    InfoLog: app.InfoLog,
    ErrorLog: app.ErrorLog,
  }
}

func (hc *HealthCheckCtx) ConfigureRouter(r chi.Router) {

   r.Route("/healthcheck", func (cr chi.Router) {

     cr.Get("/", hc.LivenessHandler)
  })
}
