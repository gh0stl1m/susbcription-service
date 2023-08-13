package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"
)

type Config struct {
  Session *scs.SessionManager
  DB *gorm.DB
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

func (app *Config) gracefulShutdown() {

  app.InfoLog.Println("Got graceful shutdown. Cleaning up pending tasks...")

  app.Wait.Wait()

  app.InfoLog.Println("Closing channels and shutting down application")
}

func (app *Config) ListenForShutdown() {

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit

  app.gracefulShutdown()
  os.Exit(0)
}

