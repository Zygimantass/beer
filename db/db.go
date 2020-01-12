package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Config *DatabaseConfig
	Connection *sql.DB
}

func New(config *DatabaseConfig) (*Database, error) {
	db := &Database {
		Config: config,
	}

	connLink:= fmt.Sprintf("%v:%v@tcp(%v)/beer", db.Config.Username,
								db.Config.Password, db.Config.URI)
	conn, err := sql.Open("mysql", connLink)
	if err != nil {
		return nil, err
	}

	db.Connection = conn

	return db, nil
}

