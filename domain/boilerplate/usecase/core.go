package usecase

import (
	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/models"
)

type boilerplateUsecase struct {
	mysqlRepository boilerplate.Repository
}

func NewBoilerplateUsecase(mysql boilerplate.Repository) boilerplate.Usecase {
	return &boilerplateUsecase{
		mysqlRepository: mysql,
	}
}

func (boilerplate boilerplateUsecase) GetAll(param map[string]interface{}) (response []models.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetAll(param)
	return
}

func (boilerplate boilerplateUsecase) GetOne(param map[string]interface{}) (response models.Boilerplate, err error) {
	response, err = boilerplate.mysqlRepository.GetOne(param)
	return
}

func (boilerplate boilerplateUsecase) Store(payload models.BoilerplatePayloadInsert) (IDs []uint64, err error) {
	var (
		data []interface{}
	)

	// PREPARE THE DATA AND INSERT TO "data"
	for _, x := range payload.Data {
		ID, _ := boilerplate.mysqlRepository.GenerateUUID()
		IDs = append(IDs, ID)
		e := []interface{}{
			ID,
			x.Column, // Custom on this line...
		}
		data = append(data, e)
	}

	err = boilerplate.mysqlRepository.Store(data)
	return
}

func (boilerplate boilerplateUsecase) Update(param map[string]interface{}, data map[string]interface{}) (err error) {
	err = boilerplate.mysqlRepository.Update(param, data)
	return
}

func (boilerplate boilerplateUsecase) Delete(param map[string]interface{}) (err error) {
	err = boilerplate.mysqlRepository.Delete(param)
	return
}
