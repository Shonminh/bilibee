package internal

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/collect"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository/api"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/logger"
	collect2 "github.com/Shonminh/bilibee/third_party/bilibili/collect"
)

type VideoCollectServiceImpl struct {
	CronTaskRepo  api.CronTaskRepo
	VideoInfoRepo api.VideoInfoRepo
	biliClient    collect2.BilibiliClient
}

func (impl *VideoCollectServiceImpl) CreateCronTask(ctx context.Context, mid int64) (err error) {
	err = db.Transaction(ctx, func(c context.Context) error {
		err = impl.CronTaskRepo.CreateCronTask(c, model.NewCronTaskTab(mid))
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

func (impl *VideoCollectServiceImpl) CollectVideoInfo(ctx context.Context) (err error) {
	cronTaskList, err := impl.CronTaskRepo.QueryUndoCronTaskList(ctx, defaultSize)
	if err != nil {
		return errors.Wrap(err, "QueryUndoCronTaskList")
	}
	for index := range cronTaskList {
		task := cronTaskList[index]
		if err = impl.doSingleTask(ctx, task); err != nil {
			return errors.Wrapf(err, "doSingleTask failed, task=%+v", task)
		}
	}
	return nil
}

// doSingleTask 针对每一个任务单次处理
func (impl *VideoCollectServiceImpl) doSingleTask(ctx context.Context, task model.CronTaskTab) (err error) {
	mid := task.GetMid()
	totalCount := 0
	defer func() {
		if err != nil {
			time.Sleep(time.Second * 3) // 停三秒
		}
		// 更新一下task的进度
		count, e := impl.VideoInfoRepo.CountVideoInfo(ctx, mid, proto.Uint32(collect.OpStatusUndo.Uint32()))
		if e != nil {
			logger.LogErrorf("CountVideoInfo failed, err=%+v", e.Error())
			return
		}
		// 更新任务列表中的total num数量和offset num数量
		e = impl.CronTaskRepo.UpdateCronTaskInfo(ctx, task.TaskId, map[string]interface{}{
			"offset_num": count,
			"total_num":  totalCount,
		})
		if e != nil {
			logger.LogErrorf("UpdateCronTaskInfo failed, err=%+v", e.Error())
			return
		}
	}()

	// 如果状态是已经完成的状态的话则不用处理了
	if task.TaskStatus == collect.TaskStatusDone.Uint32() {
		logger.LogInfof("task=%+v is done, no need to process...")
		return nil
	}
	// 从b站查询mid所有的aid list
	aidList, totalCount, err := impl.biliClient.QueryMidTotalAidList(ctx, mid, nil)
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

func (impl *VideoCollectServiceImpl) batchUpdateVideoInfo(ctx context.Context, mid int64) (err error) {
	var limit = batchSize
	needContinue := true
	for needContinue {
		videoInfoList, err := impl.VideoInfoRepo.QueryVideoInfoList(ctx, mid, nil, &limit, proto.Uint32(collect.OpStatusUndo.Uint32()))
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
			info, err := impl.biliClient.QueryVideoInfoByAid(ctx, int64(aid))
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

func (impl *VideoCollectServiceImpl) transformVideoInfo(info *collect2.VideoInfo, mid int64, aid uint64) model.VideoInfoTab {
	desc, _ := json.Marshal(info.DescV2)
	rawStr, _ := json.Marshal(info)
	infoTab := model.VideoInfoTab{
		Mid:             uint32(mid),
		Aid:             aid,
		Bvid:            info.BVID,
		Title:           info.Title,
		DescV2:          string(desc),
		Pubdate:         uint64(info.Pubdate),
		UserCtime:       uint64(info.Ctime),
		SubtitleContent: info.SubtitleContent,
		RawStr:          string(rawStr),
		OpStatus:        collect.OpStatusDone.Uint32(),
	}
	return infoTab
}
