package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/alexedwards/scs/v2"
	httpDriver "github.com/gh0stl1m/subscription-service/drivers/http"
	postgres "github.com/gh0stl1m/subscription-service/drivers/postgres"
)

type Config struct {
  Session *scs.SessionManager
  DB *sql.DB
  InfoLog *log.Logger
  ErrorLog *log.Logger
  Wait *sync.WaitGroup
}

func main() {

  db := postgres.NewConnection()
  sessionManager := httpDriver.CreateSession()
  wg := sync.WaitGroup{}

  app := Config{
    Session: sessionManager,
    DB: db,
    InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
    ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
    Wait: &wg,
  }

  app.InfoLog.Println("Server started");
}
