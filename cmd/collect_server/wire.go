//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/Shonminh/bilibee/apps/video/access/http"
	collect "github.com/Shonminh/bilibee/apps/video/wire"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/routers"
)

type CollectServerApp struct {
	router             *gin.Engine
	videoCollectSchema *http.VideoCollectHttpSchema
}

func (app *CollectServerApp) Register() {
	app.videoCollectSchema.RegisterSchema(app.router)
}

func (app *CollectServerApp) RunHttpServer(address string) error {
	return app.router.Run(address)
}

func InitCollectApp() (*CollectServerApp, error) {
	wire.Build(
		collect.CollectServerSet,
		routers.NewRouters,
		db.NewDB,
		wire.Struct(new(CollectServerApp), "*"),
	)
	return &CollectServerApp{}, nil
}
