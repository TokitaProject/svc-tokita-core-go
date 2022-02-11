package usecase

import (
	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/utils/database"
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
	var (
		data        []interface{}
		queryConfig []database.QueryConfig
	)

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

	queryInsert := boilerplate.mysqlRepository.Store(column, data)
	queryConfig = append(queryConfig, queryInsert)

	boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}

func (boilerplate boilerplateUsecase) Update(payload valueobject.BoilerplatePayloadUpdate) (err error) {
	var queryConfig []database.QueryConfig
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Param.Flag,
			},
		}
		var data = map[string]interface{}{
			"column": x.Body.Column,
		}
		queryUpdate := boilerplate.mysqlRepository.Update(param, data)
		queryConfig = append(queryConfig, queryUpdate)
	}
	boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}

func (boilerplate boilerplateUsecase) Delete(payload valueobject.BoilerplatePayloadDelete) (err error) {
	var queryConfig []database.QueryConfig
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Flag,
			},
		}
		queryDelete := boilerplate.mysqlRepository.Delete(param)
		queryConfig = append(queryConfig, queryDelete)
	}
	boilerplate.mysqlRepository.Exec(queryConfig...)
	return
}
