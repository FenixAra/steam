package steam

const (
	BaseURL = "https://api.steampowered.com"
)

type Steam struct {
}

func NewSteam() *Steam {
	return &Steam{}
}
