package config

import "os"

type Config struct {
	Token     string
	ChannelID string
}

func Load() Config {
	return Config{
		Token: os.Getenv("DISCORD_BOT_TOCKEN"),
	}
}
