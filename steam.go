package steam

const (
	BaseURL = "https://api.steampowered.com"
)

type Steam struct {
	con *Config
}

func NewSteam(key string) *Steam {
	return &Steam{
		con: NewConfig(key),
	}
}

func (s *Steam) SetSteamKey(key string) {
	s.con.Key = key
}
