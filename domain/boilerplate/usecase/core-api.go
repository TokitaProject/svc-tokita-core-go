package usecase

import (
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetAll(param)
	return
}

func (boilerplate boilerplateUsecase) GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetOne(param)
	return
}

func (boilerplate boilerplateUsecase) Store(payload valueobject.BoilerplatePayloadInsert) (err error) {
	for i := range payload.Data {
		payload.Data[i].ID, _ = boilerplate.mysqlRepository.GenerateUUID()
	}

	queryConfig, err := boilerplate.ProcessStore(payload)

	if err != nil {
		return
	}

	return boilerplate.mysqlRepository.Exec(queryConfig...)
}

func (boilerplate boilerplateUsecase) Update(payload valueobject.BoilerplatePayloadUpdate) (err error) {
	queryConfig, err := boilerplate.ProcessUpdate(payload)

	if err != nil {
		return
	}

	return boilerplate.mysqlRepository.Exec(queryConfig...)
}

func (boilerplate boilerplateUsecase) Delete(payload valueobject.BoilerplatePayloadDelete) (err error) {
	queryConfig, err := boilerplate.ProcessDelete(payload)

	if err != nil {
		return
	}

	return boilerplate.mysqlRepository.Exec(queryConfig...)
}
