package time

import "time"

func NowUint64() uint64 {
	return uint64(time.Now().Unix())
}
