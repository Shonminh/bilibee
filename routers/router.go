package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	gin2 "github.com/Shonminh/bilibee/middleware/gin"
)

func NewRouters(d *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin2.UseMysqlLogger(), gin2.UseMysql(d), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "PONG")
	})
	return router
}
