package main

import (
	"fmt"
	"github.com/krarjun90/sample-go-api/server"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	startCmd := kingpin.Command("start", "Starting app server")

	arg := kingpin.Parse()
	switch arg {
	case startCmd.FullCommand():
		server.StartApiServer()
	default:
		fmt.Printf("Command not supported : %s", arg)
	}
}
