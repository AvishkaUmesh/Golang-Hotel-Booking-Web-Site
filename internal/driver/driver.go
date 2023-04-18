package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConns = 10
const maxIdleDbConns = 5
const maxDbLifeTime = 5 * time.Minute

// ConnectDB connects to the database
func ConnectSQL(dsn string) (*DB, error) {
	db, err := newDatabase(dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetMaxIdleConns(maxIdleDbConns)
	db.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = db

	err = testDB(dbConn.SQL)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err

	}
	return nil
}
