package database

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"provar.se/webapi/lib/config"
)

var (
	cachedDB *sqlx.DB

	ErrConnectionString = errors.New("the database connection string is not valid")
)

// Setup connects to a PostgreSQL database and caches the DB connection.
// The connection string is expected to be in the following format:
//
//	postgres://<username>:<password>@<host>:<port>/<database>?sslmode=<mode>
func Setup() error {
	if cachedDB != nil {
		return nil
	}
	cfg := config.Get()
	db, err := connect(cfg.DatabaseURI)
	if err != nil {
		return err
	}
	err = ping(db)
	if err != nil {
		return err
	}
	cachedDB = db
	return nil
}

// DB returns the connected PostgreSQL database if it is available.
// NOTE: The setup function must be called before this function.
func DB() *sqlx.DB {
	if cachedDB == nil {
		log.Fatalln("Cannot use PostgreSQL database before connecting")
	}
	return cachedDB
}

// connect creates a new PostgreSQL database connection with the given connection string.
func connect(databaseURI string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ping tests the connection to a PostgreSQL database.
func ping(db *sqlx.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
