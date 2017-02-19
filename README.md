[![GoDoc](https://godoc.org/github.com/FenixAra/steam?status.svg)](https://godoc.org/github.com/FenixAra/steam)
[![Build Status](https://circleci.com/gh/FenixAra/steam.svg?&style=shield&circle-token=c211e261b26185d0ca7a2efd05b61bf39bdb177a)](https://circleci.com/gh/FenixAra/steam.svg?&style=shield&circle-token=c211e261b26185d0ca7a2efd05b61bf39bdb177a)

# steam

Steam API implementation in GoLang


    go get github.com/FenixAra/steam


Example
-------

```go
import (
	"github.com/FenixAra/steam"
)

steamAPI := steam.NewSteam(steamApiKey)
steamAPI.SetSteamKey(steamApiKey)

option := steam.NewOption(appID)
option.Count = 10
news, err := steamAPI.GetNews(option)
if err != nil {
	// Handle error
}

option := steam.NewOption(0)
option.SetSteamIDs(steamIDs)
playerSummaries, err := steamAPI.GetPlayerSummaries(option)
if err != nil {
	// Handle error
}

```
