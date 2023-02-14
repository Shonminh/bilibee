package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/collect"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/time"
)

type CronTaskRepoImpl struct {
}

const batchSize = 100

func (impl *CronTaskRepoImpl) CreateCronTask(ctx context.Context, row model.CronTaskTab) (err error) {
	now := time.NowUint64()
	row.CreateTime = now
	row.UpdateTime = now
	if err = db.GetDb(ctx).CreateInBatches(&row, batchSize).Error; err != nil {
		return errors.Wrapf(err, "CreateCronTask, row=%+v", row)
	}
	return nil
}

func (impl *CronTaskRepoImpl) QueryUndoCronTaskList(ctx context.Context, limit int) (res []model.CronTaskTab, err error) {
	err = db.GetDb(ctx).Model(&model.CronTaskTab{}).Where("task_status = ? ", collect.TaskStatusUndo).Order("update_time DESC").Limit(limit).Find(&res).Error
	return res, errors.Wrap(err, "QueryUndoCronTaskList")
}

func (impl *CronTaskRepoImpl) UpdateCronTaskInfo(ctx context.Context, taskId string, updateArgs map[string]interface{}) (err error) {
	if _, ok := updateArgs["update_time"]; !ok {
		updateArgs["update_time"] = time.NowUint64()
	}
	err = db.GetDb(ctx).Model(&model.CronTaskTab{}).Where("task_id = ? ", taskId).Updates(updateArgs).Error
	return errors.Wrapf(err, "UpdateCronTaskInfo, task_id=%s, updateArgs=%+v, err=%+v", taskId, updateArgs, err)
}

func (impl *CronTaskRepoImpl) FlushUndoStatusTask(ctx context.Context, durationSecond int) (err error) {
	updateTime := time.NowInt() - durationSecond
	err = db.GetDb(ctx).Model(&model.CronTaskTab{}).Where("update_time < ? ", updateTime).Updates(map[string]interface{}{
		"task_status": collect.TaskStatusUndo,
		"update_time": time.NowInt(),
	}).Error
	return errors.Wrap(err, "FlushCronTaskStatus")
}
