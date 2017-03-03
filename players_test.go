package steam

import (
	"os"
	"testing"
)

func TestGetPlayerSummaries(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(0)
	o.SetSteamIDs([]string{"76561197960435530"})

	playerSummaries, err := steam.GetPlayerSummaries(o)
	if err != nil {
		t.Error("Unable to get player summaries. Err: ", err)
		t.FailNow()
	}

	if playerSummaries == nil {
		t.Error("Unable to get playerSummaries. Value is nil.")
		t.FailNow()
	}

	if len(playerSummaries.Response.PlayerInfos) != 1 {
		t.Error("Unable to get playerSummaries. number of player Infos recieved is: ", len(playerSummaries.Response.PlayerInfos))
		t.FailNow()
	}
}

func TestGetFriendList(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(0)
	o.SteamID = "76561197960435530"
	o.Relationship = "friend"

	friends, err := steam.GetFriendList(o)
	if err != nil {
		t.Error("Unable to get friends list. Err: ", err)
		t.FailNow()
	}

	if friends == nil {
		t.Error("Unable to get friends. Value is nil.")
		t.FailNow()
	}

	if len(friends.List.Friends) != 302 {
		t.Error("Unable to get friends. number of player Infos recieved is: ", len(friends.List.Friends))
		t.FailNow()
	}
}

func TestGetFriendListAll(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(0)
	o.SteamID = "76561197960435530"
	o.Relationship = "all"

	friends, err := steam.GetFriendList(o)
	if err != nil {
		t.Error("Unable to get friends list. Err: ", err)
		t.FailNow()
	}

	if friends == nil {
		t.Error("Unable to get friends. Value is nil.")
		t.FailNow()
	}

	if len(friends.List.Friends) != 302 {
		t.Error("Unable to get friends. number of player Infos recieved is: ", len(friends.List.Friends))
		t.FailNow()
	}
}

func TestGetPlayerAchievements(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(440)
	o.SteamID = "76561197972495328"

	playerStats, err := steam.GetPlayerAchievements(o)
	if err != nil {
		t.Error("Unable to get player achievements. Err: ", err)
		t.FailNow()
	}

	if playerStats == nil {
		t.Error("Unable to get player achievements. Value is nil.")
		t.FailNow()
	}

	if len(playerStats.Stats.Achievements) != 520 {
		t.Error("Unable to player achievements. number of achievements Infos recieved is: ", len(playerStats.Stats.Achievements))
		t.FailNow()
	}
}

func TestGetUserStatsForGame(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(440)
	o.SteamID = "76561197972495328"

	playerStats, err := steam.GetUserStatsForGame(o)
	if err != nil {
		t.Error("Unable to get user stats. Err: ", err)
		t.FailNow()
	}

	if playerStats == nil {
		t.Error("Unable to get users stats for the game. Value is nil.")
		t.FailNow()
	}

	if len(playerStats.Stats.Stats) != 463 {
		t.Error("Unable to get users stats for the game. number of stats Infos recieved is: ", len(playerStats.Stats.Stats))
		t.FailNow()
	}
}

func TestGetOwnedGames(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(0)
	o.SetSteamID(76561197960434622, true)
	o.SetIncludeAppInfo(true, true)
	o.SetIncludePlayedFreeGames(true, true)
	o.SetAppIDs([]int{30, 40, 50, 60})

	gamesOWned, err := steam.GetOwnedGames(o)
	if err != nil {
		t.Error("Unable to get owned games. Err: ", err)
		t.FailNow()
	}

	if gamesOWned == nil {
		t.Error("Unable to get users stats for the game. Value is nil.")
		t.FailNow()
	}

	if len(gamesOWned.GamesOwnedResponse.Games) != 4 {
		t.Error("Unable to get users stats for the game. number of stats Infos recieved is: ", len(gamesOWned.GamesOwnedResponse.Games))
		t.FailNow()
	}
}
