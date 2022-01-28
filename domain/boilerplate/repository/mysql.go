package repository

import (
	"database/sql"
	"log"
	"strings"

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

	q := `
		SELECT id, uuid
		FROM boilerplate
	`

	s, p := database.QueryBuilder(param)

	query, err := db.sqlDB.Query(q+s, p...)

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
	q := `
		SELECT id, uuid
		FROM boilerplate
	`

	s, p := database.QueryBuilder(param)

	query := db.sqlDB.QueryRow(q+s, p...)

	err = query.Scan(
		&response.ID,
		&response.UUID,
	)

	return
}

func (db *mysqlBoilerplateRepository) Store(data [][]interface{}) (IDs []uint64, err error) {
	var p []interface{}

	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	q := `
		INSERT INTO boilerplate (id, column) VALUES
	`

	for _, x := range data {
		q += ` (?` + strings.Repeat(",?", len(x)) + `),`

		ID, _ := db.GenerateUUID()
		IDs = append(IDs, ID)
		p = append(p, ID)
		p = append(p, x...)
	}
	q = q[0 : len(q)-1] // TRIM THE LAST `,`

	statement, err := tx.Prepare(q)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	defer statement.Close()

	result, err := statement.Exec(p...)

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
	var p []interface{}

	tx, err := db.sqlDB.Begin()

	if err != nil {
		log.Println(err.Error())
		return
	}

	q := `
		UPDATE boilerplate SET 
	`

	for i, x := range data {
		q += i + ` = ?,`
		p = append(p, x)
	}
	q = q[0 : len(q)-1] // TRIM THE LAST `,`

	s, p2 := database.QueryBuilder(param)
	p = append(p, p2)
	statement, err := tx.Prepare(q + s)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	defer statement.Close()

	result, err := statement.Exec(p...)

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

	q := `
		DELETE FROM boilerplate
	`

	s, p := database.QueryBuilder(param)
	statement, err := tx.Prepare(q + s)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	defer statement.Close()

	result, err := statement.Exec(p...)

	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	result.RowsAffected()
	return
}
