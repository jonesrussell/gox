package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"log"

	"github.com/rivo/tview"
)

func main() {
	// Create a new Logger
	logger, err := logger.NewLogger("/tmp/gocreate.log")
	if err != nil {
		log.Fatalf("Error creating NewLogger: %v", err)
	}

	// Pass the Logger to NewWebServer
	server := webserver.NewWebServer(logger)

	// Create tview.Application and tview.Pages instances
	uiApp := tview.NewApplication()
	menuPages := tview.NewPages()

	// Pass the server, app, and pages to NewMenu
	menuInstance := ui.NewMenu(server, uiApp, menuPages)

	uiInstance := ui.NewUI()

	rootCmd := cmd.NewRootCmd(server, menuInstance, uiInstance)

	rootCmd.Execute()
}
