package main

import (
	"os"
	_boilerplateHttpDeliver "svc-boilerplate-golang/domain/boilerplate/delivery/http"
	_boilerplateRepository "svc-boilerplate-golang/domain/boilerplate/repository"
	_boilerplateUsecase "svc-boilerplate-golang/domain/boilerplate/usecase"
	"svc-boilerplate-golang/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()
	mysql := database.SetupMysqlDatabaseConnection()
	// oracle := database.SetupOracleDatabaseConnection()

	boilerplateRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
}
