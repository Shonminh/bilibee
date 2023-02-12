package repository

import (
	"context"

	"github.com/pkg/errors"

	"gorm.io/gorm/clause"

	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/time"
)

type VideoInfoRepoImpl struct {
}

func (impl *VideoInfoRepoImpl) BatchCreateVideoInfos(ctx context.Context, rows []model.VideoInfoTab) (err error) {
	now := time.NowUint64()
	for index := range rows {
		rows[index].CreateTime = now
		rows[index].UpdateTime = now
	}
	// 如果有冲突则什么都不做
	if err = db.GetDb(ctx).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(rows, batchSize).Error; err != nil {
		return errors.Wrapf(err, "BatchCreateVideoInfos, rows=%+v", rows)
	}
	return nil
}

func (impl *VideoInfoRepoImpl) QueryVideoInfoList(ctx context.Context, mid int64, aid *int64, limit *int) (res []model.VideoInfoTab, err error) {
	d := db.GetDb(ctx).Model(&model.VideoInfoTab{}).Where("mid = ? ", mid)
	if aid != nil {
		d = d.Where("aid = ? ", aid)
	}
	if limit != nil {
		d = d.Limit(*limit)
	}
	err = d.Find(&res).Error
	return res, errors.Wrapf(err, "QueryVideoInfoList, mid=%v, aid=%+v, limit=%+v", mid, aid, limit)
}
