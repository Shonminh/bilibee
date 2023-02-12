package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	RetCode RetCode     `json:"ret_code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GenCommonResponse(ctx *gin.Context, retCode RetCode, data interface{}, message string) {
	ctx.JSON(http.StatusOK, CommonResponse{
		RetCode: retCode,
		Data:    data,
		Message: message,
	})
}
