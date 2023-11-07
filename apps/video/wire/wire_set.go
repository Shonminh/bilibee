package collect

import (
	"github.com/Shonminh/bilibee/apps/video/config"
	"github.com/Shonminh/bilibee/pkg/es"
	"github.com/google/wire"

	"github.com/Shonminh/bilibee/third_party/bilibili/collect"
)

var CollectServerSet = wire.NewSet(
	NewVideoCollectHttpSchema,
	config.NewConfig,
	NewVideoCollectService,
	NewCronTaskRepo,
	NewVideoInfoRepo,
	collect.NewBilibiliClient,
	es.NewEsClient,
)

var CollectTaskSet = wire.NewSet(
	NewVideoCollectTaskSchema,
	config.NewConfig,
	NewVideoCollectService,
	NewCronTaskRepo,
	NewVideoInfoRepo,
	collect.NewBilibiliClient,
	es.NewEsClient,
)
