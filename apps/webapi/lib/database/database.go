package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

var (
	cachedDB *sql.DB

	ErrConnectionString = errors.New("the database connection string is not valid")
)

// Setup connects to a PostgreSQL database and caches the DB connection.
// The connection string is expected to be in the following format:
//
//	postgres://<username>:<password>@<host>:<port>/<database>?sslmode=<mode>
func Setup(databaseURI string) error {
	if cachedDB != nil {
		return nil
	}
	db, err := connect(databaseURI)
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
func DB() *sql.DB {
	if cachedDB == nil {
		log.Fatalln("Cannot use PostgreSQL database before connecting")
	}
	return cachedDB
}

// connect creates a new PostgreSQL database connection with the given connection string.
func connect(databaseURI string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ping tests the connection to a PostgreSQL database.
func ping(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
