package boilerplate

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

/**
why there's only one usecase interface while there can more than one repository interface?...
... because, at DDD (Domain Design Driven), there's only one set of usecase and...
... the function name inside the usecase should be unique and represent the business process...
... tl;dr: function name is telling what exactly are they doing.
*/
type Usecase interface {
	GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error)

	Store(payload valueobject.BoilerplatePayloadInsert) error
	Update(payload valueobject.BoilerplatePayloadUpdate) error
	Delete(payload valueobject.BoilerplatePayloadDelete) error

	ProcessStore(payload valueobject.BoilerplatePayloadInsert) []database.QueryConfig
	ProcessUpdate(payload valueobject.BoilerplatePayloadUpdate) []database.QueryConfig
	ProcessDelete(payload valueobject.BoilerplatePayloadDelete) []database.QueryConfig
}
