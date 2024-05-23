package interfaces

import (
	"ctfbot/internal/env"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type App interface {
	Start()
	Shutdown() error

	// Getter
	Config() *env.Config
	Handler() *handler.Mux
	Client() bot.Client

	// Commands
	CommandMention(c string) string

	// Embeds
	Footer() *discord.EmbedFooter
}
