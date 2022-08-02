package repository

import (
	"svc-boilerplate-golang/utils/database"
)

func (db *mysqlBoilerplateRepository) GenerateUUID() (uuid uint64, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *mysqlBoilerplateRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
