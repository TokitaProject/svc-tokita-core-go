package boilerplate

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

type MysqlRepository interface {
	Exec(...database.QueryConfig) error
	GenerateUUID() (uint64, error)

	GetAll(map[string]interface{}) ([]valueobject.Boilerplate, error)
	GetOne(map[string]interface{}) (valueobject.Boilerplate, error)

	Store([]string, []interface{}) database.QueryConfig
	Update(map[string]interface{}, map[string]interface{}) database.QueryConfig
	Delete(map[string]interface{}) database.QueryConfig
}

type OracleRepository interface {
	Exec(...database.QueryConfig) error

	GetAll(map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(map[string]interface{}) (response valueobject.Boilerplate, err error)

	Store([]string, []interface{}) database.QueryConfig
	Update(map[string]interface{}, map[string]interface{}) database.QueryConfig
	Delete(map[string]interface{}) database.QueryConfig
}
