package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Apps struct {
	List AppList `json:"applist"`
}

type AppList struct {
	Apps []App `json:"apps"`
}

type App struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

// Get list of all apps present in steam along with their App IDs.
func (s *Steam) GetAppList() (*Apps, error) {
	res, err := http.Get(BaseURL + "/ISteamApps/GetAppList/v2?format=json")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	apps := new(Apps)
	err = json.Unmarshal(data, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

type AppNews struct {
	News News `json:"appnews"`
}

type News struct {
	AppID    int        `json:"appid"`
	NewItems []NewsItem `json:"newsitems"`
}

type NewsItem struct {
	GID           string    `json:"gid"`
	Title         string    `json:"title"`
	URL           string    `json:"url"`
	IsExternalURL bool      `json:"is_external_url"`
	Author        string    `json:"author"`
	Contents      string    `json:"contents"`
	FeedLabel     string    `json:"feedlabel"`
	DateUnix      int       `json:"date"`
	Date          time.Time `json:"-"`
	FeedName      string    `json:"feedname"`
}

// Get news related to the application specified by application's ID.
func (s *Steam) GetNews(o *Option) (*AppNews, error) {
	res, err := http.Get(BaseURL + "/ISteamNews/GetNewsForApp/v0002?" + o.GetUrlEncode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	appNews := new(AppNews)
	err = json.Unmarshal(data, &appNews)
	if err != nil {
		return nil, err
	}

	return appNews, nil
}
