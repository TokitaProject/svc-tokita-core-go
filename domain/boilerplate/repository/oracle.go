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

func (db *oracleBoilerplateRepository) Store(column []string, data []interface{}) (err error) {
	builder := database.QueryConfig{
		TableInfo: database.TableInfo{
			TechStack: "oracle",
			Table:     "boilerplate",
			Action:    "insert",
		},
		OnInsert: database.OnInsert{
			Column: column,
			Data:   data,
		},
	}
	builder.QueryBuilder()
	builder.ExecTransaction(db.sqlDB, builder)

	return
}

func (db *oracleBoilerplateRepository) Update(param map[string]interface{}, data map[string]interface{}) (err error) {
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
	builder.ExecTransaction(db.sqlDB, builder)

	return
}

func (db *oracleBoilerplateRepository) Delete(param map[string]interface{}) (err error) {
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
	builder.ExecTransaction(db.sqlDB, builder)

	return
}
