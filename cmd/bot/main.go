package main

import (
	"ctfbot/internal/app"
	"ctfbot/internal/env"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"database/sql"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/sirupsen/logrus"
)

func main() {
	env.Load()
	c := env.Get()

	file := fmt.Sprintf("%s/db.sqlite", c.Pwd)
	u, _ := url.Parse(fmt.Sprintf("sqlite://%s", file))
	migrate := dbmate.New(u)
	migrate.SchemaFile = c.Database.SchemaFile
	migrate.MigrationsDir = []string{c.Database.MigrationsFolder}
	migrate.Log = logrus.New().Writer()

	err := migrate.CreateAndMigrate()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		logrus.Fatal(err)
	}

	boil.SetDB(db)

	app := app.New(
		app.WithConfig(c),
	)

	go app.Start()

	// Listen for CTRL+C
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	logrus.Info("Bot is now running. Press CTRL+C to exit.")
	<-done // Will block here until user hits ctrl+c

	app.Shutdown()
}
