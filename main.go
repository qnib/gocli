package main

import (
	"os"
	
	"github.com/codegangsta/cli"
)

var app *cli.App

func main() {
	// Create a new app instance
	app = cli.NewApp()
	
	// Define the app we’re building
	app.Name = "gocli"
	app.Usage = "Simple GO CLI"
	app.Version = "0.1.0"
	
	// Run the app, with the given arguments
	app.Run(os.Args)
}
