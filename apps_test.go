package steam

import "testing"

func TestGetAppList(t *testing.T) {
	steam := NewSteam("")

	apps, err := steam.GetAppList()
	if err != nil {
		t.Error("Unable to get app list. Err: ", err)
		t.FailNow()
	}

	if apps == nil {
		t.Error("Unable to get app list. Value is nil.")
		t.FailNow()
	}
}

func TestGetNews(t *testing.T) {
	steam := NewSteam("")
	o := NewOption(570)
	o.Count = 3
	o.MaxLength = 200

	appNews, err := steam.GetNews(o)
	if err != nil {
		t.Error("Unable to get app news. Err: ", err)
		t.FailNow()
	}

	if appNews == nil {
		t.Error("Unable to get app news. Value is nil.")
		t.FailNow()
	}

	if len(appNews.News.NewItems) != 3 {
		t.Error("The number of news item is not 3. Got: ", len(appNews.News.NewItems))
		t.FailNow()
	}
}

func TestGetGlobalAchievement(t *testing.T) {
	steam := NewSteam("")
	o := NewOption(730)

	achievements, err := steam.GetGlobalAchievement(o)
	if err != nil {
		t.Error("Unable to get global achievements percentage. Err: ", err)
		t.FailNow()
	}

	if achievements == nil {
		t.Error("Unable to get global achievements percentage. Value is nil.")
		t.FailNow()
	}

	if len(achievements.Percentages.Achievements) == 0 {
		t.Error("No achievements got.")
		t.FailNow()
	}
}

func TestGetGlobalStatsForGame(t *testing.T) {
	steam := NewSteam("")
	o := NewOption(17740)
	o.Names = []string{"global.map.emp_isle"}

	stats, err := steam.GetGlobalStatsForGame(o)
	if err != nil {
		t.Error("Unable to get global stats for a game and achievement. Err: ", err)
		t.FailNow()
	}

	if stats == nil {
		t.Error("Unable to get global stats for a game and achievement. Value is nil.")
		t.FailNow()
	}

	if _, ok := stats.GlobalStats.Stats["global.map.emp_isle"]; !ok {
		t.Error("Unable to get global stats for a game and achievement global.map.emp_isle.")
		t.FailNow()
	}
}
