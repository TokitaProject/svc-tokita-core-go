package usecase

import (
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) GetAllCategory(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetAllCategory(param)
	return
}
