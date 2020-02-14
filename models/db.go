package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DriverName = "mysql"
	ConnectionString = "mysql://kd13evgjkd707i0d:heo5tlc9fvcrg3kk@bbj31ma8tye2kagi.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306/ewzeh9unau2jof4l"//os.Getenv("DATABASE_URL")
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

