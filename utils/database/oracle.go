package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/godror/godror"
)

func SetupOracleDatabaseConnection() (db *sql.DB) {
	var (
		driver   = os.Getenv("ORACLE_DB_DRIVERNAME")
		username = os.Getenv("ORACLE_DB_USERNAME")
		password = os.Getenv("ORACLE_DB_PASSWORD")
		host     = os.Getenv("ORACLE_DB_HOST")
		port     = os.Getenv("ORACLE_DB_PORT")
	)

	description := fmt.Sprintf("(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%s))(CONNECT_DATA=(SERVICE_NAME=UII))))", host, port)
	connection := fmt.Sprintf("%s/%s@%s", username, password, description)

	db, err := sql.Open(driver, connection)

	if err != nil {
		log.Fatal(err)
	}

	return
}
