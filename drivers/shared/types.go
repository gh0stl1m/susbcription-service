package shared

import "github.com/go-chi/chi/v5"

type RouteConfigurer interface {
  ConfigureRouter(router chi.Router)
}

