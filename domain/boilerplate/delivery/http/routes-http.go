package http

import (
	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/domain/boilerplate"
)

type HttpBoilerplateHandler struct {
	boilerplateUsecase boilerplate.Usecase
}

func NewBoilerplateHttpHandler(boilerplate boilerplate.Usecase, httpRouter *gin.Engine) {
	handler := &HttpBoilerplateHandler{
		boilerplateUsecase: boilerplate,
	}

	// public untuk service2frontend | v1 untuk lumen, v2 untuk golang.
	public := httpRouter.Group("/public/api/v2")
	public.GET("/boilerplate", handler.GetAll)
	public.GET("/boilerplate/:uuid", handler.GetByUUID)
	public.POST("/boilerplate", handler.Store)
	public.PUT("/boilerplate", handler.Update)
	public.DELETE("/boilerplate", handler.Delete)

	// private untuk service2service | v1 untuk lumen, v2 untuk golang.
	private := httpRouter.Group("/private/api/v2")
	private.GET("/boilerplate", handler.GetAll)
	private.GET("/boilerplate/:uuid", handler.GetByUUID)
	private.POST("/boilerplate", handler.Store)
	private.PUT("/boilerplate", handler.Update)
	private.DELETE("/boilerplate", handler.Delete)
}
