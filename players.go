package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PlayerSummaries struct {
	Response PlayerSummariesResponse `json:"response"`
}

type PlayerSummariesResponse struct {
	PlayerInfos []PlayerInfo `json:"players"`
}

type PlayerInfo struct {
	SteamID                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	ProfileName              string `json:"personaname":`
	LastLogOff               int    `json:"lastlogoff"`
	ProfileURL               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`
	RealName                 string `json:"realname"`
	PrimaryClanID            string `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	PersonaStateFlags        int    `json:"personastateflags"`
	LocalCountryCode         string `json:"loccountrycode"`
	LocalStateCode           string `json:"locstatecode"`
	LocalCityID              int    `json:"loccityid"`
}

// Get player summaries using their steamids
// Options:
// SteamIDs(Mandatory) - You can set comma seperated steamIDs or use Option's method SetSteamIDs(steamIDs[]string)
// Steam should have been initialized with Steam API Key
// Or use Steam's SetSteamKey(key string) to set the Steam API key
func (s *Steam) GetPlayerSummaries(o *Option) (*PlayerSummaries, error) {
	res, err := http.Get(BaseURL + "/ISteamUser/GetPlayerSummaries/v0002?" + o.GetUrlEncode(s))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	playerSummaries := new(PlayerSummaries)
	err = json.Unmarshal(data, &playerSummaries)
	if err != nil {
		return nil, err
	}

	return playerSummaries, nil
}
