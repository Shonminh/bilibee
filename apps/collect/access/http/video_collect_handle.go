package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	http2 "github.com/Shonminh/bilibee/pkg/http"
)

type VideoCollectSchema struct {
}

func (schema *VideoCollectSchema) CreateCronTask(ctx *gin.Context) {
	// TODO implement me
	ctx.JSON(http.StatusOK, http2.CommonResponse{})
}
