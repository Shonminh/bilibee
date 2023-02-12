package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/time"
)

type VideoInfoRepoImpl struct {
}

func (impl *VideoInfoRepoImpl) CreateVideoInfo(ctx context.Context, row model.VideoInfoTab) (err error) {
	now := time.NowUint64()
	row.CreateTime = now
	row.UpdateTime = now
	if err = db.GetDb(ctx).Create(&row).Error; err != nil {
		return errors.Wrapf(err, "CreateVideoInfo, row=%+v", row)
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
