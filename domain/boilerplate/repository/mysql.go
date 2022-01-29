package repository

import (
	"database/sql"
	"log"

	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/models"
	"svc-boilerplate-golang/utils/database"
)

type mysqlBoilerplateRepository struct {
	sqlDB *sql.DB
}

func NewMysqlBoilerplateRepository(databaseConnection *sql.DB) boilerplate.Repository {
	return &mysqlBoilerplateRepository{databaseConnection}
}

func (db *mysqlBoilerplateRepository) GenerateUUID() (uuid uint64, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *mysqlBoilerplateRepository) GetAll(param map[string]interface{}) (response []models.Boilerplate, err error) {
	var result models.Boilerplate

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "mysql",
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

func (db *mysqlBoilerplateRepository) GetOne(param map[string]interface{}) (response models.Boilerplate, err error) {
	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "mysql",
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

func (db *mysqlBoilerplateRepository) Store(data []interface{}) (err error) {
	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "mysql",
			Table:     "boilerplate",
			Action:    "insert",
		},
		OnInsert: database.OnInsert{
			Column: []string{"id", "uuid"},
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

func (db *mysqlBoilerplateRepository) Update(param map[string]interface{}, data map[string]interface{}) (err error) {
	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "mysql",
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

func (db *mysqlBoilerplateRepository) Delete(param map[string]interface{}) (err error) {
	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "mysql",
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
