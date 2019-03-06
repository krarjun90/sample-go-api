package main

import (
	"fmt"
	"github.com/krarjun90/sample-go-api/config"
	"github.com/krarjun90/sample-go-api/migrations"
	"github.com/krarjun90/sample-go-api/server"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	startCmd := kingpin.Command("start", "Starting app server")
	migrateCmd := kingpin.Command("migrate", "Migrating database")
	rollbackCmd := kingpin.Command("rollback", "Rolling back database")

	config.Load()

	arg := kingpin.Parse()
	switch arg {
	case startCmd.FullCommand():
		server.StartApiServer()
	case migrateCmd.FullCommand():
		migrations.Migrate()
	case rollbackCmd.FullCommand():
		migrations.Rollback()
	default:
		fmt.Printf("Command not supported : %s", arg)
	}
}
