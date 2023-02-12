package model

import "strconv"

type CronTaskTab struct {
	Id         uint64 `gorm:"column:id;type:BIGINT(21) UNSIGNED;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	TaskId     string `gorm:"column:task_id;type:VARCHAR(64);NOT NULL"`
	TotalNum   uint32 `gorm:"column:total_num;type:INT(11) UNSIGNED;NOT NULL"`
	OffsetNum  uint32 `gorm:"column:offset_num;type:INT(11) UNSIGNED;NOT NULL"`
	TaskStatus uint32 `gorm:"column:task_status;type:INT(11) UNSIGNED;NOT NULL"`
	CreateTime uint64 `gorm:"column:create_time;type:BIGINT(21) UNSIGNED;NOT NULL"`
	UpdateTime uint64 `gorm:"column:update_time;type:BIGINT(21) UNSIGNED;NOT NULL"`
}

func (c CronTaskTab) TableName() string {
	return "cron_task_tab"
}

func genTaskId(mid int64) string {
	return strconv.FormatInt(mid, 10)
}

func (c CronTaskTab) GetMid() int64 {
	parseInt, _ := strconv.ParseInt(c.TaskId, 10, 64)
	return parseInt
}

func NewCronTaskTab(mid int64) CronTaskTab {
	res := CronTaskTab{}
	res.TaskId = genTaskId(mid)
	return res
}
