package http

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/utils/config"
	"svc-boilerplate-golang/utils/message"
	"svc-boilerplate-golang/valueobject"
)

func (handler *HttpBoilerplateHandler) GetAll(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"param":     ctx.Query("param"),
			"param_nil": nil,
			"IN": map[string][]string{
				"column_in": strings.Split(ctx.Query("column_in"), ","),
			},
			"NOT": map[string]interface{}{
				"column_not":     ctx.Query("param_not"),
				"column_not_nil": nil,
			},
			"BETWEEN": map[string]interface{}{
				"column_between": []interface{}{"A", "B"},
			},
			"LIKE": map[string]interface{}{
				"column_like": "%" + ctx.Query("param_like") + "%",
			},
			"param_special": []interface{}{">", 1},
		},
		"ORDER_BY": []string{"column ASC, column DESC"},
		"GROUP_BY": []string{"column"},
		"HAVING": [][]interface{}{
			{"column", ">", "unsigned int"},
			{"column", "<", "unsigned int"},
			{"column", "=", "int"},
		},
		"LIMIT": 100,
	}

	response, err := handler.boilerplateUsecase.GetAll(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx)
		log.Println(err)
		return
	}

	message.ReturnOk(ctx, response, param)
}

func (handler *HttpBoilerplateHandler) GetByUUID(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"uuid": ctx.Param("uuid"),
		},
	}

	response, err := handler.boilerplateUsecase.GetOne(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx)
		log.Println(err)
		return
	}

	message.ReturnOk(ctx, response, param)
}

func (handler *HttpBoilerplateHandler) Store(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadInsert

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	err = handler.boilerplateUsecase.Store(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx)
		log.Println(err)
		return
	}

	message.ReturnSuccessInsert(ctx, payload.Data)
}

func (handler *HttpBoilerplateHandler) Update(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadUpdate

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	err = handler.boilerplateUsecase.Update(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx)
		log.Println(err)
		return
	}

	message.ReturnSuccessUpdate(ctx)
}

func (handler *HttpBoilerplateHandler) Delete(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadDelete

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	err = handler.boilerplateUsecase.Delete(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx)
		log.Println(err)
		return
	}

	message.ReturnSuccessDelete(ctx)
}
