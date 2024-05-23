package settings

import (
	"context"
	"ctfbot/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Settings string

// These are the settings key that are set by commands
const (
	NotificationsChannel Settings = "notifications_channel"
	CheckInterval        Settings = "check_interval"
	TeamId               Settings = "team_id"
)

var validSettings = []Settings{
	NotificationsChannel,
	CheckInterval,
	TeamId,
}

func (s Settings) IsValid() bool {
	for _, setting := range validSettings {
		if s == setting {
			return true
		}
	}

	return false
}

var cache = make(map[Settings]string)

func SafeInit() {
	for _, settingName := range validSettings {
		s := models.Setting{
			K: string(settingName),
			V: "",
		}

		s.InsertG(context.Background(), boil.Infer())

		s.ReloadG(context.Background())

		cache[settingName] = s.V
	}
}

func Get(setting Settings) string {
	return cache[setting]
}

func Set(setting Settings, value string) error {
	_, err := models.Settings(models.SettingWhere.K.EQ(string(setting))).UpdateAllG(context.Background(), models.M{
		"v": value,
	})
	if err != nil {
		return err
	}

	cache[setting] = value

	return err
}
