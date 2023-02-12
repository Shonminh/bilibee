package routers

import (
	"net/http"

	l "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	gin2 "github.com/Shonminh/bilibee/middleware/gin"
	"github.com/Shonminh/bilibee/pkg/logger"
)

func NewRouters(d *gorm.DB) *gin.Engine {
	router := gin.New()
	output, err := logger.GetOutput()
	if err != nil {
		logger.LogPanic(err)
	}
	router.Use(l.SetLogger(l.WithLogger(func(context *gin.Context, z zerolog.Logger) zerolog.Logger { return z.Output(output).With().Logger() })))
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "PONG")
	})
	router.Use(gin2.UseMysql(d))
	return router
}
