package steam

type Config struct {
	Key string
}

func NewConfig(key string) *Config {
	return &Config{
		Key: key,
	}
}
