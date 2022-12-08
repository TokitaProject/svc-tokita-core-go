package repository

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (db *mysqlBoilerplateRepository) GetAllCategory(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	var result valueobject.Boilerplate

	builder := database.New(MYSQL, MYSQL_CATEGORY, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"category_id", "name", "last_update"},
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
			&result.Categotry,
			&result.Name,
			&result.LastUpdate,
		)

		if err != nil {
			return
		}

		response = append(response, result)
	}

	return
}
