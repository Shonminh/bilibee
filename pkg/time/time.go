package time

import "time"

func NowUint64() uint64 {
	return uint64(NowInt64())
}

func NowInt64() int64 {
	return time.Now().Unix()
}
func NowInt() int {
	return int(NowInt64())
}
