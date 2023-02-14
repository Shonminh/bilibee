package collect

import (
	"github.com/Shonminh/bilibee/apps/collect/access/crontask"
	"github.com/Shonminh/bilibee/apps/collect/access/http"
	"github.com/Shonminh/bilibee/apps/collect/api"
	"github.com/Shonminh/bilibee/apps/collect/internal"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository"
	api2 "github.com/Shonminh/bilibee/apps/collect/internal/repository/api"
	"github.com/Shonminh/bilibee/third_party/bilibili/collect"
)

func NewVideoCollectHttpSchema(VideoCollectService api.VideoCollectService) *http.VideoCollectHttpSchema {
	return &http.VideoCollectHttpSchema{
		VideoCollectService: VideoCollectService,
	}
}

func NewVideoCollectService(repo api2.CronTaskRepo, videoInfoRepo api2.VideoInfoRepo, client collect.BilibiliClient) api.VideoCollectService {
	return &internal.VideoCollectServiceImpl{CronTaskRepo: repo, VideoInfoRepo: videoInfoRepo, BiliClient: client}
}

func NewCronTaskRepo() api2.CronTaskRepo {
	return &repository.CronTaskRepoImpl{}
}

func NewVideoInfoRepo() api2.VideoInfoRepo {
	return &repository.VideoInfoRepoImpl{}
}

func NewVideoCollectTaskSchema(VideoCollectService api.VideoCollectService) *crontask.VideoCollectTaskSchema {
	return &crontask.VideoCollectTaskSchema{VideoCollectService: VideoCollectService}
}
