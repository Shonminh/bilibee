package collect

type CronTaskStatus int

const (
	TaskStatusUndo CronTaskStatus = 0
	TaskStatusDone CronTaskStatus = 1
)

func (t CronTaskStatus) Uint32() uint32 {
	return uint32(t)
}
