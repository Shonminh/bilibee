package collect

import (
	"github.com/Shonminh/bilibee/apps/collect/access/crontask"
	"github.com/Shonminh/bilibee/apps/collect/access/http"
	"github.com/Shonminh/bilibee/apps/collect/api"
	"github.com/Shonminh/bilibee/apps/collect/internal"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository"
	api2 "github.com/Shonminh/bilibee/apps/collect/internal/repository/api"
)

func NewVideoCollectHttpSchema(VideoCollectService api.VideoCollectService) *http.VideoCollectHttpSchema {
	return &http.VideoCollectHttpSchema{
		VideoCollectService: VideoCollectService,
	}
}

func NewVideoCollectService(repo api2.CronTaskRepo) api.VideoCollectService {
	return &internal.VideoCollectServiceImpl{CronTaskRepo: repo}
}

func NewCronTaskRepo() api2.CronTaskRepo {
	return &repository.CronTaskRepoImpl{}
}

func NewVideoCollectTaskSchema(VideoCollectService api.VideoCollectService) *crontask.VideoCollectTaskSchema {
	return &crontask.VideoCollectTaskSchema{VideoCollectService: VideoCollectService}
}
