package repository

import (
	"database/sql"
	"log"

	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

type oracleBoilerplateRepository struct {
	sqlDB *sql.DB
}

func NewOracleBoilerplateRepository(databaseConnection *sql.DB) boilerplate.OracleRepository {
	return &oracleBoilerplateRepository{databaseConnection}
}

func (db *oracleBoilerplateRepository) GenerateUUID() (uuid uint64, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *oracleBoilerplateRepository) GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	var result valueobject.Boilerplate

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "oracle",
			Table:     "boilerplate",
			Action:    "select",
		},
		OnSelect: database.OnSelect{
			Column: []string{"id", "uuid"},
			Where:  param,
		},
	}

	builder.QueryBuilder()

	query, err := db.sqlDB.Query(builder.Result.Query, builder.Result.Value...)

	if err != nil {
		log.Println(err.Error())
		return
	}

	defer query.Close()

	for query.Next() {
		err = query.Scan(
			&result.ID,
			&result.UUID,
		)

		if err != nil {
			log.Println(err.Error())
			return
		}

		response = append(response, result)
	}

	return
}

func (db *oracleBoilerplateRepository) GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error) {
	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "oracle",
			Table:     "boilerplate",
			Action:    "select",
		},
		OnSelect: database.OnSelect{
			Column: []string{"id", "uuid"},
			Where:  param,
		},
	}

	builder.QueryBuilder()

	query := db.sqlDB.QueryRow(builder.Result.Query, builder.Result.Value...)

	err = query.Scan(
		&response.ID,
		&response.UUID,
	)

	return
}

func (db *oracleBoilerplateRepository) Store(data []interface{}) (err error) {
	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "oracle",
			Table:     "boilerplate",
			Action:    "insert",
		},
		OnInsert: database.OnInsert{
			Column: []string{"id"},
			Data:   data,
		},
	}

	builder.QueryBuilder()

	statement, err := tx.Prepare(builder.Result.Query)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	defer statement.Close()

	result, err := statement.Exec(builder.Result.Value...)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	result.RowsAffected()
	return
}

func (db *oracleBoilerplateRepository) Update(param map[string]interface{}, data map[string]interface{}) (err error) {
	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "oracle",
			Table:     "boilerplate",
			Action:    "update",
		},
		OnUpdate: database.OnUpdate{
			Where: param,
			Data:  data,
		},
	}

	builder.QueryBuilder()

	statement, err := tx.Prepare(builder.Result.Query)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	defer statement.Close()

	result, err := statement.Exec(builder.Result.Value...)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	result.RowsAffected()
	return
}

func (db *oracleBoilerplateRepository) Delete(param map[string]interface{}) (err error) {
	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "oracle",
			Table:     "boilerplate",
			Action:    "delete",
		},
		OnDelete: database.OnDelete{
			Where: param,
		},
	}

	builder.QueryBuilder()

	statement, err := tx.Prepare(builder.Result.Query)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	defer statement.Close()

	result, err := statement.Exec(builder.Result.Value...)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	result.RowsAffected()
	return
}
