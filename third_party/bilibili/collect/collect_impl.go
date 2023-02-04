package collect

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	bg "github.com/iyear/biligo"
	"github.com/pkg/errors"
)

type BilibiliClientImpl struct {
	commonCli *bg.CommClient
}

func (impl *BilibiliClientImpl) QueryVideoInfoByAid(ctx context.Context, aid int64) (*VideoInfo, error) {
	// 查询video的信息
	videoInfo, err := impl.commonCli.VideoGetInfo(aid)
	if err != nil {
		return nil, errors.Wrap(err, "VideoGetInfo")
	}
	res := &VideoInfo{VideoInfo: videoInfo, SubtitleContent: ""}
	if videoInfo.Subtitle == nil || len(videoInfo.Subtitle.List) == 0 {
		return res, nil
	}

	// 查询字幕内容并填充
	subtitleContent, err := impl.getSubtitleContent(videoInfo)
	if err != nil {
		return nil, errors.Wrap(err, "getSubtitleContent")
	}
	res.SubtitleContent = subtitleContent
	return res, nil
}

// getSubtitleContent 获取字幕的内容并用逗号拼接成一个字符串
func (impl *BilibiliClientImpl) getSubtitleContent(videoInfo *bg.VideoInfo) (subtitleContent string, err error) {
	subtitleList := videoInfo.Subtitle.List
	set := mapset.NewSet[string]()
	for _, item := range subtitleList {
		set.Add(item.Lan)
	}
	// 只需要中文的字幕URL
	exist := false
	set.Each(func(lan string) bool { exist = checkChineseLan(lan); return exist })
	if !exist {
		return "", nil
	}

	subtitleUrl := ""
	for _, item := range subtitleList {
		if checkChineseLan(item.Lan) {
			subtitleUrl = item.SubtitleURL
			break
		}
	}
	if subtitleUrl == "" {
		return "", nil
	}

	// 请求字幕服务器http://i0.hdslb.com，获取字幕信息
	response, err := impl.commonCli.Raw(subtitleUrl, "", http.MethodGet, nil)
	if err != nil {
		return "", errors.Wrap(err, "RawParse")
	}
	subTitleData := SubTitleData{}
	if err = json.Unmarshal(response, &subTitleData); err != nil {
		return "", errors.Wrap(err, "Unmarshal")
	}

	// 拼接字幕
	var contentList []string
	for _, row := range subTitleData.Body {
		contentList = append(contentList, row.Content)
	}
	return strings.Join(contentList, ","), nil
}

type SubTitleData struct {
	FontSize        float64 `json:"font_size"`
	FontColor       string  `json:"font_color"`
	BackgroundAlpha float64 `json:"background_alpha"`
	BackgroundColor string  `json:"background_color"`
	Stroke          string  `json:"Stroke"`
	Type            string  `json:"type"`
	Lang            string  `json:"lang"`
	Version         string  `json:"version"`
	Body            []struct {
		From     float64 `json:"from"`
		To       float64 `json:"to"`
		Sid      int     `json:"sid"`
		Location int     `json:"location"`
		Content  string  `json:"content"`
	} `json:"body"`
}

func checkChineseLan(lan string) bool {
	return strings.Contains(lan, "zh")
}

func (impl *BilibiliClientImpl) QueryMidTotalAidList(ctx context.Context, mid int64) (aidList []int64, err error) {
	const size = 50 // 测试了下b站的翻页最大为50
	for index := 1; ; index++ {
		result, err := impl.commonCli.SpaceSearchVideo(mid, "", 0, "", index, size)
		if err != nil {
			return nil, errors.Wrap(err, "SpaceSearchVideo")
		}
		if result.List == nil || len(result.List.Vlist) == 0 {
			break
		}
		vlist := result.List.Vlist
		for _, v := range vlist {
			aidList = append(aidList, v.AID)
		}
		if len(result.List.Vlist) < size {
			break
		}
	}
	return aidList, nil
}