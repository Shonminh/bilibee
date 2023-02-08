package model

type VideoInfoTab struct {
	Id              uint64 `gorm:"column:id;type:BIGINT(21) UNSIGNED;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Mid             uint32 `gorm:"column:mid;type:INT(11) UNSIGNED;NOT NULL"`
	Aid             uint64 `gorm:"column:aid;type:BIGINT(21) UNSIGNED;NOT NULL"`
	Bvid            string `gorm:"column:bvid;type:VARCHAR(128);NOT NULL"`
	Title           string `gorm:"column:title;type:VARCHAR(1024);NOT NULL"`
	DescV2          string `gorm:"column:desc_v2;type:TEXT;"`
	Pubdate         uint64 `gorm:"column:pubdate;type:BIGINT(21) UNSIGNED;NOT NULL"`
	UserCtime       uint64 `gorm:"column:user_ctime;type:BIGINT(21) UNSIGNED;NOT NULL"`
	SubtitleContent string `gorm:"column:subtitle_content;type:MEDIUMTEXT;"`
	RawStr          string `gorm:"column:raw_str;type:MEDIUMTEXT;"`
	CreateTime      uint64 `gorm:"column:create_time;type:BIGINT(21) UNSIGNED;NOT NULL"`
	UpdateTime      uint64 `gorm:"column:update_time;type:BIGINT(21) UNSIGNED;NOT NULL"`
}
