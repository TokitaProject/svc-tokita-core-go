package http

import (
	"log"
	// "strings"

	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/utils/config"
	"svc-boilerplate-golang/utils/message"
	// "svc-boilerplate-golang/valueobject"
)

func (handler *HttpBoilerplateHandler) GetAllCategory(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{}}
	response, err := handler.boilerplateUsecase.GetAllCategory(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnOk(ctx, response, param)
}
