package boilerplate

import "svc-boilerplate-golang/valueobject"

type Usecase interface {
	GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error)
	Store(payload valueobject.BoilerplatePayloadInsert) (IDs []uint64, err error)
	Update(param map[string]interface{}, data map[string]interface{}) error
	Delete(param map[string]interface{}) error
}
