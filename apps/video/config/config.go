package config

type Config struct {
	ResetGetVideoTaskDurationSecond int
	ResetSyncEsTaskDurationSecond   int
}

func NewConfig() *Config {
	return &Config{
		ResetGetVideoTaskDurationSecond: 3600,
		ResetSyncEsTaskDurationSecond:   60,
	}
}
