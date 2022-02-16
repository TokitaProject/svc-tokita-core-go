package boilerplate

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

type MysqlRepository interface {
	Exec(...database.QueryConfig) error
	GenerateUUID() (uint64, error)

	GetAll(param map[string]interface{}) ([]valueobject.Boilerplate, error)
	GetOne(param map[string]interface{}) (valueobject.Boilerplate, error)

	Store(column []string, data []interface{}) database.QueryConfig
	Update(param map[string]interface{}, data map[string]interface{}) database.QueryConfig
	Delete(param map[string]interface{}) database.QueryConfig
}

type OracleRepository interface {
	Exec(...database.QueryConfig) error

	GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error)

	Store(column []string, data []interface{}) database.QueryConfig
	Update(param map[string]interface{}, data map[string]interface{}) database.QueryConfig
	Delete(param map[string]interface{}) database.QueryConfig
}
