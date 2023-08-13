package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
  Session *scs.SessionManager
  DB *sql.DB
  InfoLog *log.Logger
  ErrorLog *log.Logger
  Wait *sync.WaitGroup
}

func (app *Config) Run() {

  port := os.Getenv("SERVER_PORT")

  server := &http.Server{
    Addr: fmt.Sprintf(":%s", port),
    Handler: app.Router(),
  }

  app.InfoLog.Println("Server Running at port ", port)

  if err := server.ListenAndServe(); err != nil {

    log.Panic("Something went wrong running server", err)
  }
}

