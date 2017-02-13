package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// Get list of all apps along with their App IDs in JSON format
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
