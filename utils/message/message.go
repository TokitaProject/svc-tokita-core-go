package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func ReturnInternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Layanan sedang mengalami gangguan",
	})
}

func ReturnSuccessDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func ReturnSuccessInsert(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func ReturnSuccessUpdate(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
