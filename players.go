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
	GameID                   int    `json:"gameid"`
	GameServerIP             string `json:"gameserverip"`
	GameExtraInfo            string `json:"gameextrainfo"`
}

// Get player summaries using their steamids
//
// Options:
//
// SteamIDs(Mandatory) - You can set comma seperated steamIDs or use Option's method SetSteamIDs(steamIDs[]string)
//
// Steam should have been initialized with Steam API Key
// Or use Steam's SetSteamKey(key string) to set the Steam API key
func (s *Steam) GetPlayerSummaries(o *Option) (*PlayerSummaries, error) {
	res, err := http.Get(BaseURL + "/ISteamUser/GetPlayerSummaries/v0002?" + o.getUrlEncode(s))
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

type Friends struct {
	List FriendsList `json:"friendslist"`
}

type FriendsList struct {
	Friends []Friend `json:"friends"`
}

type Friend struct {
	SteamID      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

// Get player's friend list using their steam id
//
// Options:
//
// SteamID(Mandatory) - Steam ID of the player.
//
// Relationship(Madatory) - Relationship filter. Possible values are all, friend.
func (s *Steam) GetFriendList(o *Option) (*Friends, error) {
	res, err := http.Get(BaseURL + "/ISteamUser/GetFriendList/v0001?" + o.getUrlEncode(s))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	friends := new(Friends)
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

type PlayerStats struct {
	Stats PlayerStatistics `json:"playerstats"`
}

type PlayerStatistics struct {
	SteamID      string              `json:"steamID"`
	GameName     string              `json:"gameName"`
	Achievements []PlayerAchievement `json:"achievements"`
}

type PlayerAchievement struct {
	APIName  string `json:"apiname"`
	Achieved int    `json:"achieved"`
}

// Get player's achievements for the given app
//
// Options:
//
// SteamID(Mandatory) - Steam ID of the player you need achievement stats for.
//
// AppID(Mandatory) - The Application for which achievements are needed.
//
// l(Optional) - Language in which the achievements should be displayed.
func (s *Steam) GetPlayerAchievements(o *Option) (*PlayerStats, error) {
	res, err := http.Get(BaseURL + "/ISteamUserStats/GetPlayerAchievements/v0001?" + o.getUrlEncode(s))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	playerStats := new(PlayerStats)
	err = json.Unmarshal(data, &playerStats)
	if err != nil {
		return nil, err
	}

	return playerStats, nil
}
