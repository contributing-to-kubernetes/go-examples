package main

import (
	"log"
	"os"

	"github.com/contributing-to-kubernetes/go-examples/lesson-001-web-server-cli/app"
)

func main() {
	// This is the entrypoint into our app.
	// Here we import a cobra command and execute it. The cobra command will
	// execute some action.
	command := app.NewServerCommand()
	if err := command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
