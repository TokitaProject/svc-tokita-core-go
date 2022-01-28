package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
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

func GenerateUUID(tx *sql.DB) (uuid uint64, err error) {
	uuid = 0

	queryGetUUID := tx.QueryRow(`SELECT UUID_SHORT()`)

	err = queryGetUUID.Scan(
		&uuid,
	)

	return
}

func QueryBuilder(param map[string]interface{}) (query string, value []interface{}) {
	c := 0

	if len(param) > 0 {
		query += ` WHERE `
	}

	for i, x := range param {
		if i == "AND" {
			for g, v := range x.(map[string]interface{}) {
				if g == "IN" {
					for o, f := range v.(map[string]interface{}) {
						r := len(f.([]string)) - 1
						if r < 0 {
							r = 0
						}
						query += o + ` IN (?` + strings.Repeat(",?", r) + `) AND `
						for _, w := range f.([]string) {
							value = append(value, w)
						}
						c++
					}
				} else {
					query += g + ` = ?` + ` AND `
					value = append(value, v)
					c++
				}
			}
			query = query[0 : len(query)-4] // TRIM THE LAST `AND `
		}
	}

	return
}
