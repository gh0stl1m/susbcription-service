package postgres

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewConnection() *sql.DB {

  conn := connect()

  if conn == nil {

    log.Panic("Cannot establish connection with the DB")
  }

  return conn
}

func connect() *sql.DB {

  retry := 0

  postgresDsn := os.Getenv("POSTGRES_DSN")

  for {

    db, err := openConnection(postgresDsn)

    if err == nil {

      log.Println("Connected to the database")

      return db
    }

    if retry >= 3 { return nil }

    log.Println("Backingoff for 5 seconds")
    time.Sleep(1 * time.Second)
    retry++

    continue
  }
}

func openConnection(dsn string) (*sql.DB, error) {

  dbconn, err := gorm.Open(postgresDriver.New(postgresDriver.Config{
    DSN: dsn,
  }), &gorm.Config{})

  if err != nil {

    return nil, err
  }

  db, _ := dbconn.DB()

  if err := db.Ping(); err != nil {

    log.Println("Ping command failed")

    return nil, err
  }

  return db, nil
}


