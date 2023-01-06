package usecase

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) ProcessUpdateCategory(payload valueobject.BoilerplatePayloadUpdate) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"nis": x.Param.Nis, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"nama": x.Body.Nama, // add the data to update the row
		}

		queryUpdate, err := boilerplate.mysqlRepository.UpdateCategory(param, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessStoreCategory(payload valueobject.BoilerplatePayloadInsert) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		/**
		add data you wanted to insert on this interface{}...
		*/
		data := []interface{}{
			[]interface{}{
				x.Email,
				x.Kelas,
				x.Nis,
				x.Nama,
				// x.UserInput,
			},
		}

		/**
		column on data and this line should have same order
		*/
		column := []string{
			"email",
			"kelas",
			"nis",
			"nama",
		}

		queryInsert, err := boilerplate.mysqlRepository.StoreCategory(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (boilerplate boilerplateUsecase) ProcessDeleteCategory(payload valueobject.BoilerplatePayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"nis": x.Nis, // add the parameter to delete the row
			},
		}

		queryDelete, err := boilerplate.mysqlRepository.DeleteCategory(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
