package database

import (
	"database/sql"
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
			return err
		}

		defer statement.Close()

		_, err = statement.Exec(builder.Result.Value...)

		if err != nil {
			return err
		}
	}

	tx.Commit()

	return
}
