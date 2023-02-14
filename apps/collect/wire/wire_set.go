package collect

import "github.com/google/wire"

var CollectServerSet = wire.NewSet(
	NewVideoCollectHttpSchema,
	NewVideoCollectService,
	NewCronTaskRepo,
)

var CollectTaskSet = wire.NewSet(
	NewVideoCollectTaskSchema,
	NewVideoCollectService,
	NewCronTaskRepo,
)
