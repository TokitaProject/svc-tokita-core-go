package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupMysqlDatabaseConnection() (db *sql.DB) {
	var (
		driver   = os.Getenv("DB_DRIVERNAME")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		name     = os.Getenv("DB_NAME")
	)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, name)

	db, err := sql.Open(driver, connection)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(100 * time.Millisecond)

	return
}

func GenerateUUID(tx *sql.Tx) (uuid int64, err error) {
	uuid = 0

	queryGetUUID := tx.QueryRow(`SELECT UUID_SHORT()`)

	err = queryGetUUID.Scan(
		&uuid,
	)

	return
}