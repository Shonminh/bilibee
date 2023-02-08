package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Shonminh/bilibee/apps/collect/internal/repository/model"
)

type VideoInfoRepoImpl struct {
	Db *gorm.DB
}

func (impl *VideoInfoRepoImpl) CreateVideoInfo(ctx context.Context, row model.VideoInfoTab) (err error) {
	if err = impl.Db.Create(&row).Error; err != nil {
		return errors.Wrapf(err, "CreateVideoInfo, row=%+v", row)
	}
	return nil
}

func (impl *VideoInfoRepoImpl) QueryVideoInfoList(ctx context.Context, mid int64, aid *int64, limit *int) (res []model.VideoInfoTab, err error) {
	d := impl.Db.Model(&model.VideoInfoTab{}).Where("mid = ? ", mid)
	if aid != nil {
		d = d.Where("aid = ? ", aid)
	}
	if limit != nil {
		d = d.Limit(*limit)
	}
	err = d.Find(&res).Error
	return res, errors.Wrapf(err, "QueryVideoInfoList, mid=%v, aid=%+v, limit=%+v", mid, aid, limit)
}
