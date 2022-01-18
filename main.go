package main

import (
	"os"
	_boilerplateHttpDeliver "svc-boilerplate-golang/domain/boilerplate/delivery/http"
	_boilerplateRepo "svc-boilerplate-golang/domain/boilerplate/repository"
	_boilerplateUcase "svc-boilerplate-golang/domain/boilerplate/usecase"
	"svc-boilerplate-golang/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()
	mysql := database.SetupMysqlDatabaseConnection()

	boilerplateRepository := _boilerplateRepo.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUcase.NewBoilerplateUsecase(boilerplateRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
}
