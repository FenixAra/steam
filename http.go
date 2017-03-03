package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func makeGetRequest(url string, v interface{}) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
