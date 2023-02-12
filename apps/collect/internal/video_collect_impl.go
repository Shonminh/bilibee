package internal

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/collect/internal/repository/api"
	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
	"github.com/Shonminh/bilibee/pkg/db"
)

type VideoCollectServiceImpl struct {
	Repo api.CronTaskRepo
}

func (impl *VideoCollectServiceImpl) CreateCronTask(ctx context.Context, mid int64) (err error) {
	err = db.Transaction(ctx, func(c context.Context) error {
		err = impl.Repo.CreateCronTask(c, model.NewCronTaskTab(mid))
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
