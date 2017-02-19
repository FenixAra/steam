package steam

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Option struct {
	AppID     int
	Count     int
	MaxLength int
	Names     []string
	SteamIDs  string
}

var (
	ErrSteamIDsExceedsLimit = errors.New("Steam IDs exceeds the max allowed steam IDs")
)

func NewOption(appID int) *Option {
	return &Option{
		AppID: appID,
	}
}

func (o *Option) SetSteamIDs(ids []string) error {
	if len(ids) > 100 {
		return ErrSteamIDsExceedsLimit
	}

	o.SteamIDs = strings.Join(ids, ",")
	return nil
}

func (o *Option) GetUrlEncode(s *Steam) string {
	u := url.Values{}
	if s.con.Key != "" {
		u.Add("key", s.con.Key)
	}

	if o.SteamIDs != "" {
		u.Add("steamids", o.SteamIDs)
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

	u.Add("format", "json")
	return u.Encode()
}
