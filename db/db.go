package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // an anonymous import is only required to import the mysql driver, due to it connecting to the sql library itself
)

// Database provides a connection to the database
type Database struct {
	Config     *DatabaseConfig
	Connection *sql.DB
}

// New creates a Database object
func New(config *DatabaseConfig) (*Database, error) {
	db := &Database{
		Config: config,
	}

	connLink := fmt.Sprintf("%v:%v@tcp(%v)/beer", db.Config.Username,
		db.Config.Password, db.Config.URI)
	conn, err := sql.Open("mysql", connLink)
	if err != nil {
		return nil, err
	}

	db.Connection = conn

	return db, nil
}
