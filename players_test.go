package steam

import (
	"os"
	"testing"
)

func TestGetPlayerSummaries(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(0)
	o.SetSteamIDs([]string{"76561197960435530", "76561198146145879"})

	playerSummaries, err := steam.GetPlayerSummaries(o)
	if err != nil {
		t.Error("Unable to get player summaries. Err: ", err)
		t.FailNow()
	}

	if playerSummaries == nil {
		t.Error("Unable to get playerSummaries. Value is nil.")
		t.FailNow()
	}

	if len(playerSummaries.Response.PlayerInfos) != 2 {
		t.Error("Unable to get playerSummaries. number of player Infos recieved is: ", len(playerSummaries.Response.PlayerInfos))
		t.FailNow()
	}
}

func TestGetFriendList(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(0)
	o.SteamID = "76561198146145879"
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

	if len(friends.List.Friends) != 14 {
		t.Error("Unable to get friends. number of player Infos recieved is: ", len(friends.List.Friends))
		t.FailNow()
	}
}
