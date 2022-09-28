package repository

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (db *oracleBoilerplateRepository) GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	var result valueobject.Boilerplate

	builder := database.New(ORACLE, ORACLE_TABLE, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"id", "uuid"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query, err := db.sqlDB.Query(builder.Result.Query, builder.Result.Value...)

	if err != nil {
		return
	}

	defer query.Close()

	for query.Next() {
		err = query.Scan(
			&result.ID,
			&result.UUID,
		)

		if err != nil {
			return
		}

		response = append(response, result)
	}

	return
}

func (db *oracleBoilerplateRepository) GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error) {
	builder := database.New(ORACLE, ORACLE_TABLE, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"id", "uuid"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query := db.sqlDB.QueryRow(builder.Result.Query, builder.Result.Value...)

	err = query.Scan(
		&response.ID,
		&response.UUID,
	)

	return
}

func (db *oracleBoilerplateRepository) Store(column []string, data []interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(ORACLE, ORACLE_TABLE, INSERT)

	builder.OnInsert = database.OnInsert{
		Column: column,
		Data:   data,
	}

	err = builder.QueryBuilder()
	return
}

func (db *oracleBoilerplateRepository) Update(param map[string]interface{}, data map[string]interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(ORACLE, ORACLE_TABLE, UPDATE)

	builder.OnUpdate = database.OnUpdate{
		Where: param,
		Data:  data,
	}

	err = builder.QueryBuilder()
	return
}

func (db *oracleBoilerplateRepository) Delete(param map[string]interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(ORACLE, ORACLE_TABLE, DELETE)

	builder.OnDelete = database.OnDelete{
		Where: param,
	}

	err = builder.QueryBuilder()
	return
}
