package steam

const (
	BaseURL = "https://api.steampowered.com"
)

type Steam struct {
	conf Config
}

func NewSteam() *Steam {
	return &Steam{}
}
