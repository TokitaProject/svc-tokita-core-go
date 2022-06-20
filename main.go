package main

import (
	"log"
	"os"
	_boilerplateHttpDeliver "svc-boilerplate-golang/domain/boilerplate/delivery/http"
	_boilerplateRepository "svc-boilerplate-golang/domain/boilerplate/repository"
	_boilerplateUsecase "svc-boilerplate-golang/domain/boilerplate/usecase"
	"svc-boilerplate-golang/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()

	mysql, err := database.SetupMysqlDatabaseConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	// oracle, err := database.SetupOracleDatabaseConnection()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	boilerplateMysqlRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateMysqlRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
}
