package usecase

import (
	"svc-boilerplate-golang/domain/boilerplate"
)

/**
the struct of usecases
*/
type boilerplateUsecase struct {
	mysqlRepository boilerplate.MysqlRepository
}

/**
the initiator function for usecase
*/
func NewBoilerplateUsecase(mysql boilerplate.MysqlRepository) boilerplate.Usecase {
	return &boilerplateUsecase{
		mysqlRepository: mysql,
	}
}
