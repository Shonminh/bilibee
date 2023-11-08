package collect

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	bg "github.com/iyear/biligo"
	"github.com/pkg/errors"
)

type BilibiliClientImpl struct {
	commonCli  *bg.CommClient
	accountCli *bg.BiliClient
}

func (impl *BilibiliClientImpl) QueryVideoInfoByAid(ctx context.Context, aid int64) (*VideoInfo, error) {
	// 查询video的信息
	videoInfo, err := impl.commonCli.VideoGetInfo(aid)
	if err != nil {
		return nil, errors.Wrap(err, "VideoGetInfo")
	}
	res := &VideoInfo{VideoInfo: videoInfo, SubtitleContent: ""}
	// 查询字幕内容并填充
	subtitleContent, err := impl.getSubtitleContentV2(videoInfo)
	if err != nil {
		return nil, errors.Wrap(err, "getSubtitleContent")
	}
	res.SubtitleContent = subtitleContent
	time.Sleep(time.Millisecond * 500)
	return res, nil
}

type PlayerResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Aid      int    `json:"aid"`
		Bvid     string `json:"bvid"`
		AllowBp  bool   `json:"allow_bp"`
		NoShare  bool   `json:"no_share"`
		Cid      int    `json:"cid"`
		MaxLimit int    `json:"max_limit"`
		PageNo   int    `json:"page_no"`
		HasNext  bool   `json:"has_next"`
		IpInfo   struct {
			Ip       string `json:"ip"`
			ZoneIp   string `json:"zone_ip"`
			ZoneId   int    `json:"zone_id"`
			Country  string `json:"country"`
			Province string `json:"province"`
			City     string `json:"city"`
		} `json:"ip_info"`
		LoginMid     int    `json:"login_mid"`
		LoginMidHash string `json:"login_mid_hash"`
		IsOwner      bool   `json:"is_owner"`
		Name         string `json:"name"`
		Permission   string `json:"permission"`
		LevelInfo    struct {
			CurrentLevel int `json:"current_level"`
			CurrentMin   int `json:"current_min"`
			CurrentExp   int `json:"current_exp"`
			NextExp      int `json:"next_exp"`
			LevelUp      int `json:"level_up"`
		} `json:"level_info"`
		Vip struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgStyle               int    `json:"bg_style"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				UseImgLabel           bool   `json:"use_img_label"`
				ImgLabelUriHans       string `json:"img_label_uri_hans"`
				ImgLabelUriHant       string `json:"img_label_uri_hant"`
				ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			TvVipStatus        int    `json:"tv_vip_status"`
			TvVipPayType       int    `json:"tv_vip_pay_type"`
			TvDueDate          int    `json:"tv_due_date"`
		} `json:"vip"`
		AnswerStatus int    `json:"answer_status"`
		BlockTime    int    `json:"block_time"`
		Role         string `json:"role"`
		LastPlayTime int    `json:"last_play_time"`
		LastPlayCid  int    `json:"last_play_cid"`
		NowTime      int    `json:"now_time"`
		OnlineCount  int    `json:"online_count"`
		DmMask       struct {
			Cid     int    `json:"cid"`
			Plat    int    `json:"plat"`
			Fps     int    `json:"fps"`
			Time    int    `json:"time"`
			MaskUrl string `json:"mask_url"`
		} `json:"dm_mask"`
		NeedLoginSubtitle bool `json:"need_login_subtitle"`
		Subtitle          struct {
			AllowSubmit bool   `json:"allow_submit"`
			Lan         string `json:"lan"`
			LanDoc      string `json:"lan_doc"`
			Subtitles   []struct {
				Id          int64  `json:"id"`
				Lan         string `json:"lan"`
				LanDoc      string `json:"lan_doc"`
				IsLock      bool   `json:"is_lock"`
				SubtitleUrl string `json:"subtitle_url"`
				Type        int    `json:"type"`
				IdStr       string `json:"id_str"`
				AiType      int    `json:"ai_type"`
				AiStatus    int    `json:"ai_status"`
			} `json:"subtitles"`
		} `json:"subtitle"`
		ViewPoints      []interface{} `json:"view_points"`
		IsUgcPayPreview bool          `json:"is_ugc_pay_preview"`
		PreviewToast    string        `json:"preview_toast"`
		Options         struct {
			Is360      bool `json:"is_360"`
			WithoutVip bool `json:"without_vip"`
		} `json:"options"`
		GuideAttention []interface{} `json:"guide_attention"`
		JumpCard       []interface{} `json:"jump_card"`
		OperationCard  []interface{} `json:"operation_card"`
		OnlineSwitch   struct {
			EnableGrayDashPlayback string `json:"enable_gray_dash_playback"`
			NewBroadcast           string `json:"new_broadcast"`
			RealtimeDm             string `json:"realtime_dm"`
			SubtitleSubmitSwitch   string `json:"subtitle_submit_switch"`
		} `json:"online_switch"`
		Fawkes struct {
			ConfigVersion int `json:"config_version"`
			FfVersion     int `json:"ff_version"`
		} `json:"fawkes"`
		ShowSwitch struct {
			LongProgress bool `json:"long_progress"`
		} `json:"show_switch"`
		BgmInfo           interface{} `json:"bgm_info"`
		ToastBlock        bool        `json:"toast_block"`
		IsUpowerExclusive bool        `json:"is_upower_exclusive"`
		IsUpowerPlay      bool        `json:"is_upower_play"`
		ElecHighLevel     struct {
			PrivilegeType int    `json:"privilege_type"`
			LevelStr      string `json:"level_str"`
			Title         string `json:"title"`
			Intro         string `json:"intro"`
		} `json:"elec_high_level"`
		DisableShowUpInfo bool `json:"disable_show_up_info"`
	} `json:"data"`
}

func (impl *BilibiliClientImpl) getSubtitleContentV2(video *bg.VideoInfo) (subtitleContent string, err error) {
	raw, err := impl.accountCli.Raw("https://api.bilibili.com/x/player/v2?", "", http.MethodGet, map[string]string{
		"aid": strconv.FormatInt(video.AID, 10),
		"cid": strconv.FormatInt(video.CID, 10),
	})
	if err != nil {
		return "", errors.Wrap(err, "RawParse")
	}
	playerResponse := &PlayerResponse{}
	if err = json.Unmarshal(raw, playerResponse); err != nil {
		return "", errors.Wrap(err, "Unmarshal")
	}
	if playerResponse.Code != 0 {
		return "", errors.Errorf("playerResponse.Code != 0, %+v", playerResponse)
	}
	if len(playerResponse.Data.Subtitle.Subtitles) == 0 {
		return "", nil
	}
	var subtitleUrl = ""
	for _, item := range playerResponse.Data.Subtitle.Subtitles {
		if checkChineseLan(item.Lan) {
			subtitleUrl = item.SubtitleUrl
			break
		}
	}
	if len(subtitleUrl) == 0 {
		return "", nil
	}
	if !strings.HasPrefix(subtitleUrl, "https") {
		subtitleUrl = "https:" + subtitleUrl
	}
	return impl.curlSubtitleUrl(subtitleUrl)
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

	return impl.curlSubtitleUrl(subtitleUrl)
}

func (impl *BilibiliClientImpl) curlSubtitleUrl(subtitleUrl string) (res string, err error) {
	// 请求字幕服务器http://i0.hdslb.com，获取字幕信息
	parse, err := url.Parse(subtitleUrl)
	var payload = map[string]string{}
	for _, item := range strings.Split(parse.RawQuery, "&") {
		kv := strings.Split(item, "=")
		if len(kv) == 2 {
			payload[kv[0]] = kv[1]
		}
	}
	response, err := impl.commonCli.Raw(subtitleUrl, "", http.MethodGet, payload)
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

func (impl *BilibiliClientImpl) QueryMidTotalAidList(ctx context.Context, mid int64, limit *int64) (aidList []int64, totalCount int, err error) {
	const size = 50 // 测试了下b站的翻页最大为50
	maxSize := int64(math.MaxInt64)
	if limit != nil {
		maxSize = *limit
	}

	for index := 1; ; index++ {
		result, err := impl.commonCli.SpaceSearchVideo(mid, "", 0, "", index, size)
		if err != nil {
			return nil, 0, errors.Wrapf(err, "SpaceSearchVideo failed, mid=%+v, limit=%+v", mid, limit)
		}
		if result.List == nil || len(result.List.Vlist) == 0 {
			break
		}
		vlist := result.List.Vlist
		for _, v := range vlist {
			aidList = append(aidList, v.AID)
		}
		if totalCount == 0 {
			totalCount = result.Page.Count
		}
		if len(result.List.Vlist) < size {
			break
		}
		maxSize -= size
		if maxSize <= 0 {
			break
		}
		time.Sleep(time.Millisecond * 500) // 先限下速度，防止直接给弄崩了。
	}
	if limit != nil && int64(len(aidList)) > *limit {
		aidList = aidList[:*limit]
	}
	return aidList, totalCount, nil
}
