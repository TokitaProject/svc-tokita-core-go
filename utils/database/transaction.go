package database

import (
	"database/sql"
	"log"
)

func New(tech string, table string, command string) QueryConfig {
	return QueryConfig{
		TableInfo: TableInfo{
			TechStack: tech,
			Table:     table,
			Action:    command,
		},
	}
}

func ExecTransaction(db *sql.DB, query ...QueryConfig) (err error) {
	tx, err := db.Begin()

	if err != nil {
		return
	}

	for _, builder := range query {
		statement, err := tx.Prepare(builder.Result.Query)

		defer tx.Rollback()

		if err != nil {
			log.Println("error prepare: ", builder.Result.Query, builder.Result.Value)
			return err
		}

		defer statement.Close()

		_, err = statement.Exec(builder.Result.Value...)

		if err != nil {
			log.Println("error exec: ", builder.Result.Query, builder.Result.Value)
			return err
		}
	}

	tx.Commit()

	return
}
