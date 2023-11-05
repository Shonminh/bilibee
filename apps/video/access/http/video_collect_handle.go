package http

import (
	"github.com/Shonminh/bilibee/apps/video/internal/repository/model"
	"github.com/gin-gonic/gin"

	"github.com/Shonminh/bilibee/apps/video/api"
	"github.com/Shonminh/bilibee/pkg/http"
)

type VideoCollectHttpSchema struct {
	VideoInfoService api.VideoInfoService
}

type VideoCollectReq struct {
	Mid      int64 `json:"mid"`
	TaskType int   `json:"task_type"`
}

func (schema *VideoCollectHttpSchema) CreateCronTask(ctx *gin.Context) {
	var req VideoCollectReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		http.GenCommonResponse(ctx, http.RetCodeInvalidRequest, nil, err.Error())
		return
	}
	if err = schema.VideoInfoService.CreateCronTask(ctx.Request.Context(), req.Mid, model.TaskType(req.TaskType)); err != nil {
		http.GenCommonResponse(ctx, http.RetCodeInternalErr, nil, err.Error())
		return
	}
	http.GenCommonResponse(ctx, http.RetCodeOk, nil, "")
	return
}
