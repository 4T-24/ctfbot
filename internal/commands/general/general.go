package general

import (
	"ctfbot/internal/builder"
	"ctfbot/internal/commands/general/get"
	"ctfbot/internal/commands/general/help"
	"ctfbot/internal/commands/general/set"
	"ctfbot/internal/interfaces"

	"github.com/disgoorg/disgo/discord"
)

func NewMenu(ms *builder.MenuStore, app interfaces.App) *builder.Menu {
	return ms.NewMenu(
		builder.WithMenuName("General"),
		builder.WithEmoji(discord.Emoji{
			Name: "💻",
		}),
		builder.WithCommands(
			help.HelpCommand(ms, app),
			set.SetCommand(ms, app),
			get.GetCommand(ms, app),
		),
	)
}
