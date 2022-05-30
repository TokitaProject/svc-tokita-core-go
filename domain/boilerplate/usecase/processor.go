package usecase

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) ProcessStore(payload valueobject.BoilerplatePayloadInsert) (queryConfig []database.QueryConfig) {
	for _, x := range payload.Data {
		/**
		add data you wanted to insert on this interface{}...
		*/
		data := []interface{}{
			[]interface{}{
				x.ID,
				x.Column,
			},
		}

		/**
		column on data and this line should have same order
		*/
		column := []string{
			"id",
		}

		queryInsert := boilerplate.mysqlRepository.Store(column, data)
		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessUpdate(payload valueobject.BoilerplatePayloadUpdate) (queryConfig []database.QueryConfig) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Param.Flag, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"column": x.Body.Column, // add the data to update the row
		}
		queryUpdate := boilerplate.mysqlRepository.Update(param, data)
		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessDelete(payload valueobject.BoilerplatePayloadDelete) (queryConfig []database.QueryConfig) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"flag": x.Flag, // add the parameter to delete the row
			},
		}
		queryDelete := boilerplate.mysqlRepository.Delete(param)
		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
