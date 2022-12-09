package usecase

import (
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) GetAllCategory(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetAllCategory(param)
	return
}

func (boilerplate boilerplateUsecase) UpdateCategory(payload valueobject.BoilerplatePayloadUpdate) (err error) {
	for i := range payload.Data {
		payload.Data[i].Body.UserUpdate = payload.User
	}

	queryConfig, err := boilerplate.ProcessUpdateCategory(payload)

	if err != nil {
		return
	}

	err = boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}

func (boilerplate boilerplateUsecase) StoreCategory(payload valueobject.BoilerplatePayloadInsert) (valueobject.BoilerplatePayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].ID, _ = boilerplate.mysqlRepository.GenerateID()
		// payload.Data[i].UUID, _ = boilerplate.mysqlRepository.GenerateUUID()
		// payload.Data[i].UserInput = payload.User
	}

	queryConfig, err := boilerplate.ProcessStoreCategory(payload)

	if err != nil {
		return payload, err
	}

	return payload, boilerplate.mysqlRepository.Exec(queryConfig...)
}

func (boilerplate boilerplateUsecase) DeleteCategory(payload valueobject.BoilerplatePayloadDelete) (err error) {
	queryConfig, err := boilerplate.ProcessDeleteCategory(payload)

	if err != nil {
		return
	}

	return boilerplate.mysqlRepository.Exec(queryConfig...)
}
