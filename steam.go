package steam

const (
	BaseURL = "https://api.steampowered.com"
)

type Steam struct {
	con *Config
}

// Create a new steam API instance. You can set steam API key.
func NewSteam(key string) *Steam {
	return &Steam{
		con: NewConfig(key),
	}
}

// Set steam key for the Steam API instance
func (s *Steam) SetSteamKey(key string) {
	s.con.Key = key
}
