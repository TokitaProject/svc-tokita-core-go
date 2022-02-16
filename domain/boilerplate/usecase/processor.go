package usecase

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) ProcessStore(payload valueobject.BoilerplatePayloadInsert) (queryConfig []database.QueryConfig, IDs []uint64, err error) {
	var data []interface{}

	// Prepare the data and insert into []interface{}
	for _, x := range payload.Data {
		if x.ID == 0 {
			x.ID, _ = boilerplate.mysqlRepository.GenerateUUID()
			IDs = append(IDs, x.ID)
		}
		e := []interface{}{
			x.ID,
			x.Column, // Custom on this line...
		}
		data = append(data, e)
		column := []string{"id"}
		queryInsert := boilerplate.mysqlRepository.Store(column, data)
		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessUpdate(payload valueobject.BoilerplatePayloadUpdate) (queryConfig []database.QueryConfig, err error) {
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
	return
}

func (boilerplate boilerplateUsecase) ProcessDelete(payload valueobject.BoilerplatePayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Flag,
			},
		}
		queryDelete := boilerplate.mysqlRepository.Delete(param)
		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
