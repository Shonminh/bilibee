package video

type CronTaskStatus int

const (
	TaskStatusUndo CronTaskStatus = iota
	TaskStatusDone CronTaskStatus = 1
)

func (t CronTaskStatus) Uint32() uint32 {
	return uint32(t)
}

type OpStatus uint32

const (
	OpStatusUndo OpStatus = iota
	OpStatusDone
)

func (o OpStatus) Uint32() uint32 {
	return uint32(o)
}
