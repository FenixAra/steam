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

// Get player summaries using their steamids. Version: v0002
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

// Get player's friend list using their steam id. Version: v0001
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
	Stats        []PlayerStat        `json:"stats"`
}

type PlayerAchievement struct {
	APIName  string `json:"apiname"`
	Achieved int    `json:"achieved"`
}

type PlayerStat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Get player's achievements for the given app. Version: v0001
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

// Get list of achievements by this user for the given app/game. Version: v0002
//
// Options:
//
// SteamID(Mandatory) - Steam ID of the player you need achievement stats for.
//
// AppID(Mandatory) - The Application for which achievements are needed.
//
// l(Optional) - Language in which the achievements should be displayed.
func (s *Steam) GetUserStatsForGame(o *Option) (*PlayerStats, error) {
	res, err := http.Get(BaseURL + "/ISteamUserStats/GetUserStatsForGame/v0002?" + o.getUrlEncode(s))
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

type GamesOwned struct {
	GamesOwnedResponse GamesOwnedResponse `json:"response"`
}

type GamesOwnedResponse struct {
	GameCount int    `json:"game_count"`
	Games     []Game `json:"games"`
}

type Game struct {
	ID                       int    `json:"appid"`
	Name                     string `json:"name"`
	PlayTimeForever          int    `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
}

// Get list of games a player owns along with some playtime information, if the profile is publicly visible. Private, friends-only, and other privacy settings are not supported unless you are asking for your own personal details (ie the steam API key you are using is linked to the steamid you are requesting). Version: v0001
//
// Options:
//
// This API is a service API, you set the filters by using IsService=true.
//
// SteamID(Mandatory) - Steam ID of the player you need achievement stats for.
//
// IncludeAppInfo(Optional) - Include game name and logo information in the output. The default is to return appids only.
//
// IncludePlayedFreeGames(Optional) - By default, free games like Team Fortress 2 are excluded (as technically everyone owns them). If include_played_free_games is set, they will be returned if the player has played them at some point. This is the same behavior as the games list on the Steam Community.
//
// AppIDs(Optional) - App IDs to verify if the user owns it.
func (s *Steam) GetOwnedGames(o *Option) (*GamesOwned, error) {
	res, err := http.Get(BaseURL + "/IPlayerService/GetOwnedGames/v0001?" + o.getUrlEncode(s))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	gamesOWned := new(GamesOwned)
	err = json.Unmarshal(data, &gamesOWned)
	if err != nil {
		return nil, err
	}

	return gamesOWned, nil
}
