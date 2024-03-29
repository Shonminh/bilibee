package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shonminh/bilibee/apps/video"
	"github.com/Shonminh/bilibee/apps/video/config"
	time2 "github.com/Shonminh/bilibee/pkg/time"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/video/internal/repository/api"
	"github.com/Shonminh/bilibee/apps/video/internal/repository/model"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/logger"
	collect2 "github.com/Shonminh/bilibee/third_party/bilibili/collect"
)

type VideoInfoServiceImpl struct {
	CronTaskRepo  api.CronTaskRepo
	VideoInfoRepo api.VideoInfoRepo
	BiliClient    collect2.BiliBiliClient
	Config        *config.Config
	EsClient      *elasticsearch8.Client
}

func (impl *VideoInfoServiceImpl) CreateCronTask(ctx context.Context, mid int64, taskType model.TaskType) (err error) {
	err = db.Transaction(ctx, func(c context.Context) error {
		err = impl.CronTaskRepo.CreateCronTask(c, model.NewCronTaskTab(mid, taskType))
		if err != nil {
			if !db.IsMysqlDuplicateErr(err) {
				return errors.Wrap(err, "CreateCronTask")
			}
			// 是重复键冲突的话也返回正常。
		}
		return nil
	})
	return err
}

const defaultSize int = 100

func (impl *VideoInfoServiceImpl) CollectVideoInfo(ctx context.Context) (err error) {
	cronTaskList, err := impl.CronTaskRepo.QueryUndoCronTaskList(ctx, defaultSize, model.TaskTypeGetVideoInfo)
	if err != nil {
		return errors.Wrap(err, "QueryUndoCronTaskList")
	}
	for index := range cronTaskList {
		task := cronTaskList[index]
		if err = impl.doSingleTask(ctx, task); err != nil {
			return errors.Wrapf(err, "doSingleTask failed, task=%+v", task)
		}
	}
	logger.LogInfo("CollectVideoInfo process...")
	return nil
}

// doSingleTask 针对每一个任务单次处理
func (impl *VideoInfoServiceImpl) doSingleTask(ctx context.Context, task model.CronTaskTab) (err error) {
	mid := task.GetMid()
	totalCount := 0
	defer func() {
		if err != nil {
			time.Sleep(time.Second * 3) // 停三秒
		}
		// 更新一下task的进度
		count, e := impl.VideoInfoRepo.CountVideoInfo(ctx, mid, proto.Uint32(video.OpStatusDone.Uint32()))
		if e != nil {
			logger.LogErrorf("CountVideoInfo failed, err=%+v", e.Error())
			return
		}
		// 更新任务列表中的total num数量和offset num数量
		updateArgs := map[string]interface{}{"offset_num": count, "total_num": totalCount, "task_status": video.TaskStatusDone.Uint32()}
		e = impl.CronTaskRepo.UpdateCronTaskInfo(ctx, task.TaskId, updateArgs)
		if e != nil {
			logger.LogErrorf("UpdateCronTaskInfo failed, err=%+v", e.Error())
			return
		}
	}()

	// 如果状态是已经完成的状态的话则不用处理了
	if task.TaskStatus == video.TaskStatusDone.Uint32() {
		logger.LogInfof("task=%+v is done, no need to process...")
		return nil
	}
	// 从b站查询mid所有的aid list
	aidList, totalCount, err := impl.BiliClient.QueryMidTotalAidList(ctx, mid, nil)
	if err != nil {
		return errors.Wrap(err, "QueryMidTotalAidList")
	}

	// 先存储对应的aid list到video_info_tab表中
	if err = impl.VideoInfoRepo.BatchCreateVideoInfos(ctx, genVideInfoTab(aidList, mid)); err != nil {
		return errors.Wrap(err, "BatchCreateVideoInfos")
	}

	// 查询需要更新的vide_info信息。
	if err = impl.batchUpdateVideoInfo(ctx, mid); err != nil {
		return errors.Wrap(err, "batchUpdateVideoInfo")
	}
	return nil
}

func genVideInfoTab(aidList []int64, mid int64) (rows []model.VideoInfoTab) {
	rows = make([]model.VideoInfoTab, len(aidList), len(aidList))
	for index := range aidList {
		rows[index] = model.VideoInfoTab{Mid: uint32(mid), Aid: uint64(aidList[index])}
	}
	return rows
}

const batchSize = 100

func (impl *VideoInfoServiceImpl) batchUpdateVideoInfo(ctx context.Context, mid int64) (err error) {
	var limit = batchSize
	needContinue := true
	for needContinue {
		videoInfoList, err := impl.VideoInfoRepo.QueryVideoInfoList(ctx, mid, nil, &limit, proto.Uint32(video.OpStatusUndo.Uint32()))
		if err != nil {
			return errors.Wrap(err, "QueryVideoInfoList")
		}
		if len(videoInfoList) == 0 { // 为0的话则说明没有数据了
			break
		}
		if len(videoInfoList) < limit { // 终结条件是查询的次数少于limit的时候，就可以
			needContinue = false
		}
		for index := range videoInfoList {
			aid := videoInfoList[index].Aid
			info, err := impl.BiliClient.QueryVideoInfoByAid(ctx, int64(aid))
			if err != nil {
				logger.LogErrorf("QueryVideoInfoByAid failed, aid=%s, info=%+v", aid, info)
				continue
			}
			// 把b站的数据转换成数据库的model
			infoTab := impl.transformVideoInfo(info, mid, aid)
			// 更新数据库video_info_tab
			if err = impl.VideoInfoRepo.UpdateVideoInfo(ctx, infoTab); err != nil {
				logger.LogErrorf("UpdateVideoInfo failed, aid=%s, need_update_info=%+v", aid, infoTab)
				continue
			}
		}
	}
	return nil
}

const videoUri = "https://www.bilibili.com/video/"

func (impl *VideoInfoServiceImpl) transformVideoInfo(info *collect2.VideoInfo, mid int64, aid uint64) model.VideoInfoTab {
	desc, _ := json.Marshal(info.DescV2)
	rawStr, _ := json.Marshal(info)
	url := ""
	if len(info.BVID) > 0 {
		url = videoUri + info.BVID
	}
	infoTab := model.VideoInfoTab{
		Mid:             uint32(mid),
		Aid:             aid,
		Bvid:            info.BVID,
		Url:             url,
		Title:           info.Title,
		DescV2:          string(desc),
		Pubdate:         uint64(info.Pubdate),
		UserCtime:       uint64(info.Ctime),
		SubtitleContent: info.SubtitleContent,
		RawStr:          string(rawStr),
		OpStatus:        video.OpStatusDone.Uint32(),
	}
	return infoTab
}

func (impl *VideoInfoServiceImpl) ResetTaskUndoStatus(ctx context.Context) (err error) {
	// 刷新获取视频信息的任务
	if err = impl.CronTaskRepo.FlushUndoStatusTask(ctx, impl.Config.ResetGetVideoTaskDurationSecond, model.TaskTypeGetVideoInfo); err != nil {
		return errors.Wrap(err, "FlushUndoStatusTask")
	}
	// 刷新同步es的任务
	if err = impl.CronTaskRepo.FlushUndoStatusTask(ctx, impl.Config.ResetSyncEsTaskDurationSecond, model.TaskTypeSyncVideoInfoToEs); err != nil {
		return errors.Wrap(err, "FlushUndoStatusTask")
	}
	return nil
}

func (impl *VideoInfoServiceImpl) SyncVideoInfoToEs(ctx context.Context) (err error) {
	cronTaskList, err := impl.CronTaskRepo.QueryUndoCronTaskList(ctx, defaultSize, model.TaskTypeSyncVideoInfoToEs)
	if err != nil {
		return errors.Wrap(err, "QueryUndoCronTaskList")
	}

	// 如果索引不存的话则进行创建。
	if err = impl.existVideoInfoIndex(ctx); err != nil {
		if errors.Cause(err) != ErrIndexNotExist {
			return errors.Wrap(err, "existVideoInfoIndex")
		}
		if err = impl.createVideoInfoIndex(ctx); err != nil {
			return errors.Wrap(err, "createVideoInfoIndex")
		}
	}

	// 每个任务单独处理
	for index := range cronTaskList {
		task := cronTaskList[index]
		if err = impl.doSingleSyncTask(ctx, task); err != nil {
			return errors.Wrapf(err, "doSingleSyncTask failed, task=%+v", task)
		}
	}
	return nil
}

func (impl *VideoInfoServiceImpl) doSingleSyncTask(ctx context.Context, task model.CronTaskTab) error {
	if task.TaskStatus == video.TaskStatusDone.Uint32() {
		logger.LogInfof("task=%+v is done, no need to process...")
		return nil
	}

	mid := task.GetMid()
	defer func() {
		totalCount, err := impl.VideoInfoRepo.CountVideoInfo(ctx, mid, nil)
		if err != nil {
			logger.LogErrorf("CountVideoInfo failed, err=%+v", err.Error())
			return
		}
		successCount, err := impl.VideoInfoRepo.CountVideoInfo(ctx, mid, proto.Uint32(video.OpStatusDone.Uint32()))
		if err != nil {
			logger.LogErrorf("CountVideoInfo failed, err=%+v", err.Error())
			return
		}
		err = impl.CronTaskRepo.UpdateCronTaskInfo(ctx, task.TaskId, map[string]interface{}{
			"offset_num":  successCount,
			"total_num":   totalCount,
			"task_status": video.TaskStatusDone.Uint32(),
		})
		if err != nil {
			logger.LogErrorf("UpdateCronTaskInfo failed, err=%+v", err.Error())
			return
		}
	}()

	rows, err := impl.VideoInfoRepo.QueryVideoInfosByUpdateTime(ctx, mid, impl.Config.ScanVideoInfoUpdateTimeDurationSecond)
	if err != nil {
		return errors.Wrap(err, "QueryVideoInfosByUpdateTime")
	}
	if len(rows) == 0 {
		return nil
	}
	// 批量同步到es中
	if err = impl.batchSyncVideoInfoToEs(ctx, rows); err != nil {
		return errors.Wrap(err, "batchSyncVideoInfoToEs")
	}
	return nil
}

func (impl *VideoInfoServiceImpl) batchSyncVideoInfoToEs(ctx context.Context, rows []*model.VideoInfoTab) error {

	for index := 0; index < len(rows); index += batchSize {
		start := index
		end := start + batchSize
		if end > len(rows) {
			end = len(rows)
		}
		subRows := rows[start:end]
		// 批量同步到es中
		if err := impl.syncVideoInfoToEs(ctx, subRows); err != nil {
			return errors.Wrap(err, "syncVideoInfoToEs")
		}
	}
	return nil
}

var ErrIndexNotExist = errors.New("index not exist")

func (impl *VideoInfoServiceImpl) existVideoInfoIndex(ctx context.Context) error {
	indicesExists := impl.EsClient.Indices.Exists
	exists, err := indicesExists([]string{model.VideoInfoMapping{}.IndexName()}, indicesExists.WithContext(ctx), indicesExists.WithPretty())
	if err != nil {
		return errors.Wrap(err, "Indices.Exists")
	}
	if exists == nil || exists.StatusCode == http.StatusOK {
		return nil
	}
	return ErrIndexNotExist
}

func (impl *VideoInfoServiceImpl) createVideoInfoIndex(ctx context.Context) error {
	create := impl.EsClient.Indices.Create
	response, err := create(model.VideoInfoMapping{}.IndexName(), create.WithBody(strings.NewReader(model.VideoInfoMapping{}.Mapping())), create.WithContext(ctx), create.WithPretty())
	if err != nil {
		return errors.Wrap(err, "Indices.Create")
	}
	if response == nil || response.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New(fmt.Sprintf("Indices.Create failed, response=%+v", response))
}

func (impl *VideoInfoServiceImpl) syncVideoInfoToEs(ctx context.Context, rows []*model.VideoInfoTab) error {
	if len(rows) == 0 {
		return nil
	}

	esModels := make([]*model.VideoInfoMapping, len(rows))
	for i := range rows {
		row := rows[i]
		esModels[i] = &model.VideoInfoMapping{
			Id:              row.Id,
			Mid:             row.Mid,
			Aid:             row.Aid,
			Bvid:            row.Bvid,
			Url:             row.Url,
			Title:           row.Title,
			DescV2:          row.DescV2,
			Pubdate:         row.Pubdate,
			UserCtime:       row.UserCtime,
			SubtitleContent: row.SubtitleContent,
			OpStatus:        row.OpStatus,
			CreateTime:      row.CreateTime,
			UpdateTime:      row.UpdateTime,
			EsUpdateTime:    time2.NowUint64(),
		}
	}
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{Client: impl.EsClient})
	if err != nil {
		return errors.Wrap(err, "NewBulkIndexer")
	}

	defer func() {
		e := bulkIndexer.Close(ctx)
		if e != nil {
			logger.LogErrorf("bulkIndexer Close failed, err=%+v", e)
		}
	}()

	for i := range esModels {
		esModel := esModels[i]
		marshal, _ := json.Marshal(esModel)
		err = bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Index:      esModel.IndexName(),
			Action:     "index",
			DocumentID: strconv.FormatUint(esModel.Id, 10),
			Body:       bytes.NewReader(marshal),
		})
		if err != nil {
			logger.LogErrorf("bulkIndexer Add failed, err=%+v", err)
		}
	}

	return nil
}
