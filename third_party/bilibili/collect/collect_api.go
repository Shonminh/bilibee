package collect

import (
	"context"

	bg "github.com/iyear/biligo"
)

type VideoInfo struct {
	*bg.VideoInfo
	SubtitleContent string `json:"subtitle_content"`
}

type BiliBiliClient interface {
	// QueryVideoInfoByAid 根据aid查询视频信息
	QueryVideoInfoByAid(ctx context.Context, aid int64) (*VideoInfo, error)
	// QueryMidTotalAidList 获取UP主的投稿的所有视频aid列表
	// mid up主的id
	QueryMidTotalAidList(ctx context.Context, mid int64, limit *int64) (aidList []int64, totalCount int, err error)
}
