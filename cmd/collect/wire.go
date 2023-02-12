//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/Shonminh/bilibee/apps/collect/access/http"
	collect "github.com/Shonminh/bilibee/apps/collect/wire"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/routers"
)

type CollectApp struct {
	router             *gin.Engine
	videoCollectSchema *http.VideoCollectSchema
	db                 *gorm.DB
}

func (app *CollectApp) Register() {
	app.videoCollectSchema.RegisterSchema(app.router)
}

func (app *CollectApp) RunHttpServer(address string) error {
	return app.router.Run(address)
}

func InitCollectApp() (*CollectApp, error) {
	wire.Build(
		collect.CollectSet,
		routers.NewRouters,
		db.NewDB,
		wire.Struct(new(CollectApp), "*"),
	)
	return &CollectApp{}, nil
}
