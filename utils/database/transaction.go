package database

import (
	"database/sql"
	"log"
)

func ExecTransaction(db *sql.DB, query ...QueryConfig) (err error) {
	tx, err := db.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, builder := range query {
		statement, err := tx.Prepare(builder.Result.Query)

		defer tx.Rollback()

		if err != nil {
			log.Println(err.Error())
			break
		}

		defer statement.Close()

		_, err = statement.Exec(builder.Result.Value...)

		if err != nil {
			log.Println(err.Error())
			break
		}
	}

	if err == nil {
		tx.Commit()
	}

	return
}
