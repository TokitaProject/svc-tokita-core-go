package boilerplate

import "svc-boilerplate-golang/valueobject"

type MysqlRepository interface {
	GenerateUUID() (uuid uint64, err error)
	GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error)
	Store(column []string, data []interface{}) error
	Update(param map[string]interface{}, data map[string]interface{}) error
	Delete(param map[string]interface{}) error
}

type OracleRepository interface {
	GenerateUUID() (uuid uint64, err error)
	GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error)
	Store(column []string, data []interface{}) error
	Update(param map[string]interface{}, data map[string]interface{}) error
	Delete(param map[string]interface{}) error
}
