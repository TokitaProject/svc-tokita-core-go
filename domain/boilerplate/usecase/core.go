package usecase

import (
	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/valueobject"
)

type boilerplateUsecase struct {
	mysqlRepository boilerplate.MysqlRepository
}

func NewBoilerplateUsecase(mysql boilerplate.MysqlRepository) boilerplate.Usecase {
	return &boilerplateUsecase{
		mysqlRepository: mysql,
	}
}

func (boilerplate boilerplateUsecase) GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetAll(param)
	return
}

func (boilerplate boilerplateUsecase) GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetOne(param)
	return
}

func (boilerplate boilerplateUsecase) Store(payload valueobject.BoilerplatePayloadInsert) (IDs []uint64, err error) {
	var data []interface{}

	// Prepare the data and insert into []interface{}
	for _, x := range payload.Data {
		ID, _ := boilerplate.mysqlRepository.GenerateUUID()
		IDs = append(IDs, ID)
		e := []interface{}{
			ID,
			x.Column, // Custom on this line...
		}
		data = append(data, e)
	}
	column := []string{"id"}

	err = boilerplate.mysqlRepository.Store(column, data)
	return
}

func (boilerplate boilerplateUsecase) Update(payload valueobject.BoilerplatePayloadUpdate) (err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Param.Flag,
			},
		}
		var data = map[string]interface{}{
			"column": x.Body.Column,
		}
		err = boilerplate.mysqlRepository.Update(param, data)
	}
	return
}

func (boilerplate boilerplateUsecase) Delete(payload valueobject.BoilerplatePayloadDelete) (err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Flag,
			},
		}
		err = boilerplate.mysqlRepository.Delete(param)
	}
	return
}
