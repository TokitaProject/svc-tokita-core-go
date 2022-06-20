package database

import (
	"database/sql"
	"fmt"
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
			return fmt.Errorf("query: %s | value: %s | message: %s", builder.Result.Query, builder.Result.Value, err.Error())
		}

		defer statement.Close()

		_, err = statement.Exec(builder.Result.Value...)

		if err != nil {
			return fmt.Errorf("query: %s | value: %s | message: %s", builder.Result.Query, builder.Result.Value, err.Error())
		}
	}

	tx.Commit()

	return
}
