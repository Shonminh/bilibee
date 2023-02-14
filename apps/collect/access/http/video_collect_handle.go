package http

import (
	"github.com/gin-gonic/gin"

	"github.com/Shonminh/bilibee/apps/collect/api"
	"github.com/Shonminh/bilibee/pkg/http"
)

type VideoCollectHttpSchema struct {
	VideoCollectService api.VideoCollectService
}

type VideoCollectReq struct {
	Mid int64 `json:"mid"`
}

func (schema *VideoCollectHttpSchema) CreateCronTask(ctx *gin.Context) {
	var req VideoCollectReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		http.GenCommonResponse(ctx, http.RetCodeInvalidRequest, nil, err.Error())
		return
	}
	if err = schema.VideoCollectService.CreateCronTask(ctx.Request.Context(), req.Mid); err != nil {
		http.GenCommonResponse(ctx, http.RetCodeInternalErr, nil, err.Error())
		return
	}
	http.GenCommonResponse(ctx, http.RetCodeOk, nil, "")
	return
}
