package steam

import "testing"

func TestGetAppList(t *testing.T) {
	steam := NewSteam()

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
	steam := NewSteam()
	option := NewOption(570)
	option.Count = 3
	option.MaxLength = 200

	appNews, err := steam.GetNews(option)
	if err != nil {
		t.Error("Unable to get app news. Err: ", err)
		t.FailNow()
	}

	if appNews == nil {
		t.Error("Unable to get app news. Value is nil.")
		t.FailNow()
	}

	if len(appNews.News.NewItems) != 3 {
		t.Error("The number of news item is no 3.")
		t.FailNow()
	}
}
