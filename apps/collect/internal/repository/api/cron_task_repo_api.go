package api

import (
	"context"

	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
)

type CronTaskRepo interface {
	CreateCronTask(ctx context.Context, row model.CronTaskTab) (err error)
	QueryUndoCronTaskList(ctx context.Context, limit int) (res []model.CronTaskTab, err error)
	UpdateCronTaskInfo(ctx context.Context, taskId string, updateArgs map[string]interface{}) (err error)
}
