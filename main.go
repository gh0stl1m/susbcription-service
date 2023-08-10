package main

import (

  postgres "github.com/gh0stl1m/subscription-service/drivers/postgres"
)

func main() {

  db := postgres.NewConnection()
  db.Ping()
}
