package collect

import (
	"github.com/google/wire"

	"github.com/Shonminh/bilibee/third_party/bilibili/collect"
)

var CollectServerSet = wire.NewSet(
	NewVideoCollectHttpSchema,
	NewVideoCollectService,
	NewCronTaskRepo,
	NewVideoInfoRepo,
	collect.NewBilibiliClient,
)

var CollectTaskSet = wire.NewSet(
	NewVideoCollectTaskSchema,
	NewVideoCollectService,
	NewCronTaskRepo,
	NewVideoInfoRepo,
	collect.NewBilibiliClient,
)
