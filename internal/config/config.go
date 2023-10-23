package config

type Config struct {
	Token     string
	ChannelID string
}

func Load() Config {
	return Config{
		Token: "PASTE BOT TOKEN HERE",
	}
}
