package usecase

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) ProcessStore(payload valueobject.BoilerplatePayloadInsert) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		/**
		add data you wanted to insert on this interface{}...
		*/
		data := []interface{}{
			[]interface{}{
				x.ID,
				x.UUID,
				x.UserInput,
			},
		}

		/**
		column on data and this line should have same order
		*/
		column := []string{
			"id", "uuid", "user_input",
		}

		queryInsert, err := boilerplate.mysqlRepository.Store(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessUpdate(payload valueobject.BoilerplatePayloadUpdate) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.Param.UUID, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"user_update": x.Body.UserUpdate, // add the data to update the row
		}

		queryUpdate, err := boilerplate.mysqlRepository.Update(param, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessDelete(payload valueobject.BoilerplatePayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.UUID, // add the parameter to delete the row
			},
		}

		queryDelete, err := boilerplate.mysqlRepository.Delete(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
