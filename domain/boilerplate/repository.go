package boilerplate

import "svc-boilerplate-golang/models"

type Repository interface {
	GenerateUUID() (uint64, error)
	GetIn(where string, in []interface{}) ([]models.Boilerplate, error)
	GetAll(param map[string]interface{}) ([]models.Boilerplate, error)
	GetOne(param map[string]interface{}) (models.Boilerplate, error)
	Store(data [][]interface{}) ([]uint64, error)
	Update(param map[string]interface{}, data map[string]interface{}) error
	Delete(param map[string]interface{}) error
}
