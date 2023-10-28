package config

type CommandConfig struct {
	MeowCount int
}

var currentCommandConfig = &CommandConfig{
	MeowCount: 1,
}

func LoadCommandConfig() *CommandConfig {
	return currentCommandConfig
}

func UpdateMeowCount(count int) {
	currentCommandConfig.MeowCount = count
}
