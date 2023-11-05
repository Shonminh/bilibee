package model

type VideoInfoMapping struct {
	Id              uint64 `json:"id"`
	Mid             uint32 `json:"mid"`
	Aid             uint64 `json:"aid"`
	Bvid            string `json:"bvid"`
	Url             string `json:"url"`
	Title           string `json:"title"`
	DescV2          string `json:"desc_v2"`
	Pubdate         uint64 `json:"pubdate"`
	UserCtime       uint64 `json:"user_ctime"`
	SubtitleContent string `json:"subtitle_content"`
	RawStr          string `json:"raw_str"`
	OpStatus        uint32 `json:"op_status"`
	CreateTime      uint64 `json:"create_time"`
	UpdateTime      uint64 `json:"update_time"`
}
