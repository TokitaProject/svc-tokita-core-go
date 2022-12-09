package http

import (
	"log"
	// "strings"

	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/utils/config"
	"svc-boilerplate-golang/utils/message"
	"svc-boilerplate-golang/valueobject"
)

func (handler *HttpBoilerplateHandler) GetAllCategory(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{}}
	response, err := handler.boilerplateUsecase.GetAllCategory(param)

	if err != nil {
		if err.Error() == config.ERROR_BIND_JSON {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnOk(ctx, response, param)
}

func (handler *HttpBoilerplateHandler) UpdateCategory(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadUpdate

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	payload.User = ctx.Request.Header.Get("X-Member")

	err = handler.boilerplateUsecase.UpdateCategory(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessUpdate(ctx)
}

func (handler *HttpBoilerplateHandler) StoreCategory(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadInsert

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	payload.User = ctx.Request.Header.Get("X-Member")

	feedback, err := handler.boilerplateUsecase.StoreCategory(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessInsert(ctx, feedback.Data)
}

func (handler *HttpBoilerplateHandler) DeleteCategory(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadDelete

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	err = handler.boilerplateUsecase.DeleteCategory(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessDelete(ctx)
}
