package collect

import "github.com/google/wire"

var CollectSet = wire.NewSet(
	NewVideoCollectSchema,
	NewVideoCollectService,
	NewCronTaskRepo,
)
