package collect

import (
	"github.com/Shonminh/bilibee/apps/collect/access/http"
	"github.com/Shonminh/bilibee/apps/collect/api"
	"github.com/Shonminh/bilibee/apps/collect/internal"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository"
	api2 "github.com/Shonminh/bilibee/apps/collect/internal/repository/api"
)

func NewVideoCollectSchema(VideoCollectService api.VideoCollectService) *http.VideoCollectSchema {
	return &http.VideoCollectSchema{
		VideoCollectService: VideoCollectService,
	}
}

func NewVideoCollectService(repo api2.CronTaskRepo) api.VideoCollectService {
	return &internal.VideoCollectServiceImpl{CronTaskRepo: repo}
}

func NewCronTaskRepo() api2.CronTaskRepo {
	return &repository.CronTaskRepoImpl{}
}
