package commands

import (
	"ctfbot/internal/builder"
	"ctfbot/internal/commands/general"
	"ctfbot/internal/interfaces"

	"github.com/sirupsen/logrus"
)

func RegisterCommands(app interfaces.App) *builder.MenuStore {
	ms := builder.NewMenuStore(app)

	general.NewMenu(ms, app)

	createCommands, err := ms.RegisterCommands()
	if err != nil {
		logrus.Fatal("Couldn't load commands")
	}

	_, err = app.Client().Rest().SetGlobalCommands(
		app.Client().ApplicationID(),
		createCommands,
	)
	if err != nil {
		logrus.WithError(err).Error("couldn't register commands")
	}

	return ms
}
