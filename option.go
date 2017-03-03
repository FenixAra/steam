package steam

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Option struct {
	AppID                  int
	Count                  int
	MaxLength              int
	Names                  []string
	SteamIDs               string
	SteamID                string
	Relationship           string
	Language               string
	IncludeAppInfo         bool
	IncludePlayedFreeGames bool
	HasJSONData            bool
	JSONData               JSONData
}

type JSONData struct {
	SteamID                int   `json:"steamid"`
	AppIDs                 []int `json:"appids_filter"`
	IncludeAppInfo         bool  `json:"include_appinfo"`
	IncludePlayedFreeGames bool  `json:"include_played_free_games"`
}

var (
	ErrSteamIDsExceedsLimit = errors.New("Steam IDs exceeds the max allowed steam IDs")
)

// Get a new Option instance. You can set an App ID to Option.
func NewOption(appID int) *Option {
	return &Option{
		AppID: appID,
		JSONData: JSONData{
			AppIDs: []int{},
		},
	}
}

// Set steamIDs to Option
func (o *Option) SetSteamIDs(ids []string) error {
	if len(ids) > 100 {
		return ErrSteamIDsExceedsLimit
	}

	o.SteamIDs = strings.Join(ids, ",")
	return nil
}

// Set APP IDs JSON filters
func (o *Option) SetAppIDs(appIDs []int) {
	o.JSONData.AppIDs = appIDs
	o.HasJSONData = true
}

// Set the steam ID filter, If IsService is set true then the filters are only sent as JSON input
func (o *Option) SetSteamID(steamID int, IsService bool) {
	if IsService {
		o.JSONData.SteamID = steamID
		o.HasJSONData = true
	} else {
		o.SteamID = strconv.Itoa(steamID)
	}
}

func (o *Option) SetIncludeAppInfo(includeAppInfo bool, IsService bool) {
	if IsService {
		o.JSONData.IncludeAppInfo = includeAppInfo
		o.HasJSONData = true
	} else {
		o.IncludeAppInfo = includeAppInfo
	}
}

func (o *Option) SetIncludePlayedFreeGames(includePlayedFreeGames bool, IsService bool) {
	if IsService {
		o.JSONData.IncludePlayedFreeGames = includePlayedFreeGames
		o.HasJSONData = true
	} else {
		o.IncludePlayedFreeGames = includePlayedFreeGames
	}
}

func (o *Option) getUrlEncode(s *Steam) string {
	u := url.Values{}

	u.Add("format", "json")

	if s.con.Key != "" {
		u.Add("key", s.con.Key)
	}

	if o.HasJSONData {
		content, _ := json.Marshal(o.JSONData)
		u.Add("input_json", string(content))
		return u.Encode()
	}

	if o.SteamIDs != "" {
		u.Add("steamids", o.SteamIDs)
	}

	if o.SteamID != "" {
		u.Add("steamid", o.SteamID)
	}

	if o.Language != "" {
		u.Add("l", o.Language)
	}

	if o.Relationship != "" {
		u.Add("relationship", o.Relationship)
	}

	if o.AppID > 0 {
		u.Add("appid", strconv.Itoa(o.AppID))
		u.Add("gameid", strconv.Itoa(o.AppID))
	}

	if o.Count > 0 {
		u.Add("count", strconv.Itoa(o.Count))
	}

	if o.MaxLength > 0 {
		u.Add("maxlength", strconv.Itoa(o.MaxLength))
	}

	if len(o.Names) > 0 {
		for i, name := range o.Names {
			u.Add("name["+strconv.Itoa(i)+"]", name)
		}
		u.Add("count", strconv.Itoa(len(o.Names)))
	}

	if o.IncludeAppInfo {
		u.Add("include_appinfo", strconv.FormatBool(o.IncludeAppInfo))
	}

	if o.IncludePlayedFreeGames {
		u.Add("include_played_free_games", strconv.FormatBool(o.IncludePlayedFreeGames))
	}

	return u.Encode()
}
