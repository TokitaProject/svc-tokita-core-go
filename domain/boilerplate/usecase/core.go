package usecase

import (
	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/models"
)

type boilerplateUsecase struct {
	boilerplateRepository boilerplate.Repository
}

func NewBoilerplateUsecase(boilerplate boilerplate.Repository) boilerplate.Usecase {
	return &boilerplateUsecase{
		boilerplateRepository: boilerplate,
	}
}

func (boilerplate boilerplateUsecase) GetAll(param map[string]interface{}) (response []models.Boilerplate, err error) {
	response, err = boilerplate.boilerplateRepository.GetAll(param)
	return
}

func (boilerplate boilerplateUsecase) GetByUUID(param map[string]interface{}) (response models.Boilerplate, err error) {
	response, err = boilerplate.boilerplateRepository.GetOne(param)
	return
}

func (boilerplate boilerplateUsecase) Store(payload models.BoilerplatePayloadInsert) (err error) {
	var data [][]interface{}

	// PREPARE THE DATA AND INSERT TO "data"
	for _, x := range payload.Data {
		e := []interface{}{
			x.Column, // Custom on this line...
		}
		data = append(data, e)
	}

	_, err = boilerplate.boilerplateRepository.Store(data)
	return
}

func (boilerplate boilerplateUsecase) Update(param map[string]interface{}, data map[string]interface{}) (err error) {
	err = boilerplate.boilerplateRepository.Update(param, data)
	return
}

func (boilerplate boilerplateUsecase) Delete(param map[string]interface{}) (err error) {
	err = boilerplate.boilerplateRepository.Delete(param)
	return
}
