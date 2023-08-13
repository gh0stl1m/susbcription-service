package main

import (
	"log"
	"os"
	"sync"

	httpDriver "github.com/gh0stl1m/subscription-service/drivers/http"
	postgres "github.com/gh0stl1m/subscription-service/drivers/postgres"
)

func main() {

  db := postgres.NewConnection()
  sessionManager := httpDriver.CreateSession()
  wg := sync.WaitGroup{}

  app := httpDriver.Config{
    Session: sessionManager,
    DB: db,
    InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
    ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
    Wait: &wg,
  }

  go app.ListenForShutdown()
  
  app.Run()
}
