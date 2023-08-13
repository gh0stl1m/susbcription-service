package http

import "net/http"

func (app *Config) LoadSession(next http.Handler) http.Handler {

  return app.Session.LoadAndSave(next)
}
