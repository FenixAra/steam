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
