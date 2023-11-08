package collect

import (
	"github.com/Shonminh/bilibee/apps/video/access/crontask"
	"github.com/Shonminh/bilibee/apps/video/access/http"
	"github.com/Shonminh/bilibee/apps/video/api"
	"github.com/Shonminh/bilibee/apps/video/config"
	"github.com/Shonminh/bilibee/apps/video/internal"
	"github.com/Shonminh/bilibee/apps/video/internal/repository"
	api2 "github.com/Shonminh/bilibee/apps/video/internal/repository/api"
	"github.com/Shonminh/bilibee/third_party/bilibili/collect"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

func NewVideoCollectHttpSchema(VideoCollectService api.VideoInfoService) *http.VideoCollectHttpSchema {
	return &http.VideoCollectHttpSchema{
		VideoInfoService: VideoCollectService,
	}
}

func NewVideoCollectService(repo api2.CronTaskRepo, videoInfoRepo api2.VideoInfoRepo, client collect.BiliBiliClient,
	config *config.Config, esClient *elasticsearch8.Client) api.VideoInfoService {
	return &internal.VideoInfoServiceImpl{CronTaskRepo: repo, VideoInfoRepo: videoInfoRepo, BiliClient: client,
		Config: config, EsClient: esClient}
}

func NewCronTaskRepo() api2.CronTaskRepo {
	return &repository.CronTaskRepoImpl{}
}

func NewVideoInfoRepo() api2.VideoInfoRepo {
	return &repository.VideoInfoRepoImpl{}
}

func NewVideoCollectTaskSchema(VideoCollectService api.VideoInfoService) *crontask.VideoCollectTaskSchema {
	return &crontask.VideoCollectTaskSchema{VideoCollectService: VideoCollectService}
}
