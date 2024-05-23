package get

import (
	"ctfbot/internal/builder"
	"ctfbot/internal/env"
	"ctfbot/internal/interfaces"
	"ctfbot/internal/settings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

const (
	commandName        = "get"
	commandDescription = "Get a setting's value"
)

type getCommand struct {
	c   *env.Config
	app interfaces.App
	ms  *builder.MenuStore
}

func GetCommand(ms *builder.MenuStore, app interfaces.App) *builder.Command {
	var cmd getCommand

	cmd.app = app
	cmd.ms = ms
	cmd.c = app.Config()

	return builder.NewCommand(
		builder.WithCommandName(commandName),
		builder.WithDescription(commandDescription),
		builder.WithRegisterFunc(func(h *handler.Mux) error {
			h.Command("/"+commandName, cmd.HandleCommand)
			return nil
		}),
		builder.WithSlashCommand(discord.SlashCommandCreate{
			Name:        commandName,
			Description: commandDescription,
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "key",
					Description: "The setting's key",
					Required:    true,
				},
			},
		}),
	)
}

func (cmd *getCommand) HandleCommand(e *handler.CommandEvent) error {
	key := e.SlashCommandInteractionData().String("key")

	s := settings.Settings(key)
	if !s.IsValid() {
		return e.Respond(
			discord.InteractionResponseTypeCreateMessage,
			discord.NewMessageCreateBuilder().
				SetContent("Invalid setting key!").
				SetEphemeral(true),
		)
	}

	v := settings.Get(s)

	return e.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetContentf("Here it is : `%v`", v).
			SetEphemeral(true),
	)
}
