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

func (boilerplate boilerplateUsecase) Store(payload valueobject.BoilerplatePayloadInsert) (IDs []uint64, err error) {
	queryConfig, IDs, err := boilerplate.ProcessStore(payload)

	if err != nil {
		return
	}

	err = boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}

func (boilerplate boilerplateUsecase) Update(payload valueobject.BoilerplatePayloadUpdate) (err error) {
	queryConfig, err := boilerplate.ProcessUpdate(payload)

	if err != nil {
		return
	}

	err = boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}

func (boilerplate boilerplateUsecase) Delete(payload valueobject.BoilerplatePayloadDelete) (err error) {
	queryConfig, err := boilerplate.ProcessDelete(payload)

	if err != nil {
		return
	}

	err = boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}
