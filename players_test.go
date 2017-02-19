package steam

import (
	"os"
	"testing"
)

func TestGetPlayerSummaries(t *testing.T) {
	steam := NewSteam(os.Getenv("STEAM_KEY"))
	o := NewOption(17740)
	o.SetSteamIDs([]string{"76561197960435530", "76561198146145879"})

	playerSummaries, err := steam.GetPlayerSummaries(o)
	if err != nil {
		t.Error("Unable to get global stats for a game and achievement. Err: ", err)
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
