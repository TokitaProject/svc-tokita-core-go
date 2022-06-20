package main

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"log"
	"os"

	"svc-boilerplate-golang/utils/database"

	_boilerplateHttpDeliver "svc-boilerplate-golang/domain/boilerplate/delivery/http"
	_boilerplateRepository "svc-boilerplate-golang/domain/boilerplate/repository"
	_boilerplateUsecase "svc-boilerplate-golang/domain/boilerplate/usecase"
)

func main() {
	routers := gin.Default()

	mysql := ConnectMySQL()
	//orcl:= ConnectOracle()

	boilerplateMysqlRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateMysqlRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
}

func ConnectMySQL() (mysql *sql.DB) {
	mysql, err := database.SetupMysqlDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}

func ConnectOracle() (oracle *sql.DB) {
	oracle, err := database.SetupOracleDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
