package steam

type Config struct {
	key string
}

// Get a new steam config
func NewConfig(key string) *Config {
	conf := &Config{}
	conf.key = key
	return conf
}
