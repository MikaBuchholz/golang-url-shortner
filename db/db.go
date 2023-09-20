package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		fmt.Println("Failed to open DB with dsn: ", dsn, " and error: ", err)
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	err = testDB(db)

	if err != nil {
		return nil, err
	}

	dbConn.DB = db
	return dbConn, nil

}

func testDB(d *sql.DB) error {
	err := d.Ping()

	if err != nil {
		fmt.Println("Failed to ping DB, error: ", err)
		return err
	}

	fmt.Println("DB connection working")
	return nil
}
