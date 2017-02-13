package steam

import (
	"net/url"
	"strconv"
)

type Option struct {
	AppID     int
	Count     int
	MaxLength int
}

func NewOption(appID int) *Option {
	return &Option{
		AppID: appID,
	}
}

func (o *Option) GetUrlEncode() string {
	u := url.Values{}
	if o.AppID > 0 {
		u.Add("appid", strconv.Itoa(o.AppID))
	}

	if o.Count > 0 {
		u.Add("count", strconv.Itoa(o.Count))
	}

	if o.MaxLength > 0 {
		u.Add("maxlength", strconv.Itoa(o.MaxLength))
	}

	u.Add("format", "json")
	return u.Encode()
}
