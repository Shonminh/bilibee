package http

import "github.com/gin-gonic/gin"

func (schema *VideoCollectHttpSchema) RegisterSchema(router *gin.Engine) {
	router.POST("/api/cron_task/create", schema.CreateCronTask)
}
