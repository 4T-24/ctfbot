# CTF Bot

## Development

Quick setup :

### DBMate

To install DBMate :

```
sudo curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
sudo chmod +x /usr/local/bin/dbmate
```

You can then use it in vscode (thanks to the autofilled env) :

```
dbmate up
```

This should generate `db.sqlite`.

### SQLBoiler

You first need to copy `sqlboiler.toml.example` into `sqlboiler.toml`.
To install SQLBoiler :

```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
```

To generate (with filled up `db.sqlite`) :

```
go generate
```

### Env

An example environement is available in `.env.example`, you juste need to fill in the token and ID of your bot (findable in discord's developer website)

###

You can then run the bot using `go run cmd/bot/main.go`