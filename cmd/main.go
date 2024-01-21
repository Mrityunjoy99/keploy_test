package main

import (
	"fmt"
	"os"

	"github.com/sample_project/cmd/cmdopts"
	"github.com/sample_project/migration"
	"github.com/sample_project/src/infrastructure/appserver"
	"github.com/sample_project/src/infrastructure/database/seed"
)

func main() {
	args := cmdopts.Parse()

	const MaxArgLength = 2

	if len(os.Args) < MaxArgLength {
		fmt.Println("Invalid command")
		os.Exit(1)
	}

	command := os.Args[len(os.Args)-1]

	switch command {
	case "migrate":
		migration.Migrate(args)
	case "start":
		appserver.Start(args)
	case "seed":
		seed.SeedDB()
	default:
		panic("Invalid command")
	}
}
