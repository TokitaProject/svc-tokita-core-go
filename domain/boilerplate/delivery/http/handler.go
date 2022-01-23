package http

import (
	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/domain/boilerplate"
	"svc-boilerplate-golang/models"
	"svc-boilerplate-golang/utils/message"
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

func (handler *HttpBoilerplateHandler) GetAll(ctx *gin.Context) {
	var param = map[string]interface{}{
		"param": ctx.Query("param"),
	}

	response, err := handler.boilerplateUsecase.GetAll(param)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx)
		return
	}

	message.ReturnOk(ctx, response, param)
}

func (handler *HttpBoilerplateHandler) GetByUUID(ctx *gin.Context) {
	var param = map[string]interface{}{
		"uuid": ctx.Param("uuid"),
	}

	response, err := handler.boilerplateUsecase.GetOne(param)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx)
		return
	}

	message.ReturnOk(ctx, response, param)
}

func (handler *HttpBoilerplateHandler) Store(ctx *gin.Context) {
	var payload models.BoilerplatePayloadInsert

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err, payload)
		return
	}

	err = handler.boilerplateUsecase.Store(payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err, payload)
		return
	}

	message.ReturnSuccessInsert(ctx, payload)
}

func (handler *HttpBoilerplateHandler) Update(ctx *gin.Context) {
	var payload models.BoilerplatePayloadUpdate

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err, payload)
		return
	}

	var param = map[string]interface{}{
		"flag": payload.Param.Flag,
	}

	var data = map[string]interface{}{
		"column": payload.Data.Column,
	}

	err = handler.boilerplateUsecase.Update(param, data)

	if err != nil {
		message.ReturnBadRequest(ctx, err, param)
		return
	}

	message.ReturnOk(ctx, data, param)
}

func (handler *HttpBoilerplateHandler) Delete(ctx *gin.Context) {
	var payload models.BoilerplatePayloadDetele

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err, payload)
		return
	}

	var param = map[string]interface{}{
		"flag": payload.Flag,
	}

	err = handler.boilerplateUsecase.Delete(param)

	if err != nil {
		message.ReturnBadRequest(ctx, err, param)
		return
	}

	message.ReturnSuccessDelete(ctx)
}
