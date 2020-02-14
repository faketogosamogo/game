package models

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)

var (
	DriverName = "postgres"
	ConnectionString = os.Getenv("DATABASE_URL")
)

type DB struct {
	*sql.DB
}

func NewDB(driver, conn string)(*DB, error){
	db, err := sql.Open(driver, conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

