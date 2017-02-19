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
//
// Options:
//
// AppID(Mandatory) - App for which you need the news.
//
// Count(Optional) - Number of news articles you want to receive.
//
// MaxLength(Optional) - Maximum length of each news entry.
func (s *Steam) GetNews(o *Option) (*AppNews, error) {
	res, err := http.Get(BaseURL + "/ISteamNews/GetNewsForApp/v0002?" + o.getUrlEncode(s))
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

type GlobalAchievements struct {
	Percentages AchievementPercentages `json:"achievementpercentages"`
}

type AchievementPercentages struct {
	Achievements []Achievement `json:"achievements"`
}

type Achievement struct {
	Name       string  `json:"name"`
	Percentage float32 `json:"percent"`
}

// Get global acheivements overview of a specific game in percentage
//
// Options:
//
// AppID(Mandatory) - App for which you need the global achievements.
func (s *Steam) GetGlobalAchievement(o *Option) (*GlobalAchievements, error) {
	res, err := http.Get(BaseURL + "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002?" + o.getUrlEncode(s))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	acheivements := new(GlobalAchievements)
	err = json.Unmarshal(data, &acheivements)
	if err != nil {
		return nil, err
	}

	return acheivements, nil
}

type GlobalStatsResponse struct {
	GlobalStats GlobalStats `json:"response"`
}

type GlobalStats struct {
	Stats  map[string]Stat `json:"globalstats"`
	Result int             `json:"result"`
}

type Stat struct {
	Total string `json:"total"`
}

// Get global stats detail for a particular achievement for the given game
//
// Options:
//
// AppID(Mandatory) - App for which you need the global stats.
//
// Names(Mandatory, atleast 1) - The achievement name/names for which you need the global stats.
func (s *Steam) GetGlobalStatsForGame(o *Option) (*GlobalStatsResponse, error) {
	res, err := http.Get(BaseURL + "/ISteamUserStats/GetGlobalStatsForGame/v0001?" + o.getUrlEncode(s))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	stats := new(GlobalStatsResponse)
	err = json.Unmarshal(data, &stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
