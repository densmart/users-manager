package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, message)
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func EmptyResponse(ctx *gin.Context) {
	ctx.String(http.StatusNoContent, "")
}
