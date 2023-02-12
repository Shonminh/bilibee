package collect

type CronTaskStatus int

const (
	TaskStatusUndo CronTaskStatus = iota
	TaskStatusDone CronTaskStatus = 1
)

func (t CronTaskStatus) Uint32() uint32 {
	return uint32(t)
}

type OpStatus int

const (
	OpStatusUndo OpStatus = iota
	OpStatusDone
)
