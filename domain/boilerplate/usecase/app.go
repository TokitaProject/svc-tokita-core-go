package usecase

import (
	"svc-boilerplate-golang/domain/boilerplate"
)

type boilerplateUsecase struct {
	mysqlRepository boilerplate.MysqlRepository
}

func NewBoilerplateUsecase(mysql boilerplate.MysqlRepository) boilerplate.Usecase {
	return &boilerplateUsecase{
		mysqlRepository: mysql,
	}
}
