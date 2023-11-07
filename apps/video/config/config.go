package config

type Config struct {
	ResetGetVideoTaskDurationSecond       int
	ResetSyncEsTaskDurationSecond         int
	ScanVideoInfoUpdateTimeDurationSecond int
}

func NewConfig() *Config {
	return &Config{
		ResetGetVideoTaskDurationSecond:       3600,
		ResetSyncEsTaskDurationSecond:         60,
		ScanVideoInfoUpdateTimeDurationSecond: 3600 * 24 * 365, // 每次同步1年之内的记录，后面可以改下这个值。
	}
}
