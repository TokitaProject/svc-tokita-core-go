package boilerplate

import "svc-boilerplate-golang/models"

type Usecase interface {
	GetAll(param map[string]interface{}) (response []models.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response models.Boilerplate, err error)
	Store(payload models.BoilerplatePayloadInsert) (IDs []uint64, err error)
	Update(param map[string]interface{}, data map[string]interface{}) error
	Delete(param map[string]interface{}) error
}
