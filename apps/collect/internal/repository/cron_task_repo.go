package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Shonminh/bilibee/apps/collect"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
)

type CronTaskRepoImpl struct {
	Db *gorm.DB
}

const batchSize = 100

func (impl *CronTaskRepoImpl) CreateCronTask(ctx context.Context, row model.CronTaskTab) (err error) {
	if err = impl.Db.CreateInBatches(&row, batchSize).Error; err != nil {
		return errors.Wrapf(err, "CreateCronTask, row=%+v", row)
	}
	return nil
}

func (impl *CronTaskRepoImpl) QueryUndoCronTaskList(ctx context.Context, limit int) (res []model.CronTaskTab, err error) {
	err = impl.Db.Model(&model.CronTaskTab{}).Where("task_status = ? ", collect.TaskStatusUndo).Order("update_time DESC").Limit(limit).Find(&res).Error
	return res, errors.Wrap(err, "QueryUndoCronTaskList")
}

func (impl *CronTaskRepoImpl) UpdateCronTaskInfo(ctx context.Context, taskId string, updateArgs map[string]interface{}) (err error) {
	err = impl.Db.Model(&model.CronTaskTab{}).Where("task_id = ? ", taskId).Updates(updateArgs).Error
	return errors.Wrapf(err, "UpdateCronTaskInfo, task_id=%s, updateArgs=%+v, err=%+v", taskId, updateArgs, err)
}
