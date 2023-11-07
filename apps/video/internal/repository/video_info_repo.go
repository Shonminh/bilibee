package repository

import (
	"context"

	"github.com/pkg/errors"

	"gorm.io/gorm/clause"

	"github.com/Shonminh/bilibee/apps/video/internal/repository/model"
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

func (impl *VideoInfoRepoImpl) QueryVideoInfoList(ctx context.Context, mid int64, aid *int64, limit *int, opStatus *uint32) (res []model.VideoInfoTab, err error) {
	d := db.GetDb(ctx).Model(&model.VideoInfoTab{}).Where("mid = ? ", mid)
	if aid != nil {
		d = d.Where("aid = ? ", aid)
	}
	if opStatus != nil {
		d = d.Where("op_status = ?", opStatus)
	}

	if limit != nil {
		d = d.Limit(*limit)
	}
	err = d.Find(&res).Error
	return res, errors.Wrapf(err, "QueryVideoInfoList, mid=%v, aid=%+v, limit=%+v", mid, aid, limit)
}

func (impl *VideoInfoRepoImpl) UpdateVideoInfo(ctx context.Context, row model.VideoInfoTab) (err error) {
	err = db.GetDb(ctx).Table(row.TableName()).Where("aid = ?", row.Aid).Omit("id", "aid", "create_time").Updates(&row).Error
	if err != nil {
		return errors.Wrapf(err, "UpdateVideoInfo failed, row=%+v", row)
	}
	return nil
}

func (impl *VideoInfoRepoImpl) CountVideoInfo(ctx context.Context, mid int64, opStatus *uint32) (res int64, err error) {
	d := db.GetDb(ctx).Model(&model.VideoInfoTab{}).Where("mid = ? ", mid)
	if opStatus != nil {
		d = d.Where("op_status = ?", opStatus)
	}
	if err = d.Count(&res).Error; err != nil {
		return 0, errors.Wrapf(err, "CountVideoInfo, mid=%+v, op_status=%+v", mid, opStatus)
	}
	return res, nil
}

func (impl *VideoInfoRepoImpl) QueryVideoInfosByUpdateTime(ctx context.Context, mid int64, updateTimeDuration int) (res []*model.VideoInfoTab, err error) {
	now := time.NowInt()
	err = db.GetDb(ctx).Model(&model.VideoInfoTab{}).Where("mid = ? AND update_time >= ? AND update_time < ? ", mid, now-updateTimeDuration, now).Order("update_time ASC").Find(&res).Error
	return res, errors.Wrapf(err, "QueryVideoInfosByUpdateTime, mid=%+v, updateTimeDuration=%+v", mid, updateTimeDuration)
}
