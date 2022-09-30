package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const ERROR_STATUS = "Layanan sedang mengalami gangguan"
const SUCCESS_UPDATE_STATUS = "Berhasil melakukan pembaharuan data"

func ReturnOk(ctx *gin.Context, data interface{}, param interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":  data,
		"param": param,
	})
}

func ReturnBadRequest(ctx *gin.Context, message interface{}, param interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"param":   param,
	})
}

func ReturnInternalServerError(ctx *gin.Context, msg ...string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message":               ERROR_STATUS,
		"message for developer": msg,
	})
}

func ReturnSuccessInsert(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func ReturnSuccessDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func ReturnSuccessUpdate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": SUCCESS_UPDATE_STATUS,
	})
}
