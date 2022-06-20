package main

import (
	"database/sql"
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

	mysql := ConnectMySQL()
	// oracle := ConnectOracle()

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
