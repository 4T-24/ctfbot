package env

import (
	"ctfbot/internal/values"

	"github.com/disgoorg/snowflake/v2"
)

type Config struct {
	Mode values.Mode `env:"MODE" envDefault:"unset"`
	Pwd  string      `env:"PWD" envDefault:"."`

	// Database
	Database struct {
		SchemaFile       string `env:"SCHEMA_FILE" envDefault:"/sql/schema.sql"`
		MigrationsFolder string `env:"MIGRATIONS_FOLDER" envDefault:"/sql/migrations"`
	} `envPrefix:"DB_"`

	// Discord
	Discord struct {
		Token string       `env:"TOKEN" envDefault:""`
		AppID snowflake.ID `env:"APP_ID" envDefault:""`
	} `envPrefix:"DISCORD_"`

	// App parameters
	App struct {
		BotColor int `env:"BOT_COLOR" envDefault:"3859607"`
	} `envPrefix:"APP_"`
}

func Get() *Config {
	return &cfg
}
