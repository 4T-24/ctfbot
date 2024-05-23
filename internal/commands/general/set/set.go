package set

import (
	"ctfbot/internal/builder"
	"ctfbot/internal/env"
	"ctfbot/internal/interfaces"
	"ctfbot/internal/settings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

const (
	commandName        = "set"
	commandDescription = "Set a setting's value"
)

type setCommand struct {
	c   *env.Config
	app interfaces.App
	ms  *builder.MenuStore
}

func SetCommand(ms *builder.MenuStore, app interfaces.App) *builder.Command {
	var cmd setCommand

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
				discord.ApplicationCommandOptionString{
					Name:        "val",
					Description: "The setting's value",
					Required:    true,
				},
			},
		}),
	)
}

func (cmd *setCommand) HandleCommand(e *handler.CommandEvent) error {
	key := e.SlashCommandInteractionData().String("key")
	val := e.SlashCommandInteractionData().String("val")

	s := settings.Settings(key)
	if !s.IsValid() {
		return e.Respond(
			discord.InteractionResponseTypeCreateMessage,
			discord.NewMessageCreateBuilder().
				SetContent("Invalid setting key!").
				SetEphemeral(true),
		)
	}

	if err := settings.Set(s, val); err != nil {
		return e.Respond(
			discord.InteractionResponseTypeCreateMessage,
			discord.NewMessageCreateBuilder().
				SetContent("Something went wrong!").
				SetEphemeral(true),
		)
	}

	return e.Respond(
		discord.InteractionResponseTypeCreateMessage,
		discord.NewMessageCreateBuilder().
			SetContent("All good!").
			SetEphemeral(true),
	)
}
