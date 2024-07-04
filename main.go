package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"log"

	"github.com/rivo/tview"
)

func main() {
	// Create a new Logger
	logger := logger.NewLogger("/tmp/gocreate.log")
	err := logger.Init()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	// Pass the Logger to NewServer
	server := webserver.NewServer(logger)

	// Create tview.Application and tview.Pages instances
	uiApp := tview.NewApplication()
	menuPages := tview.NewPages()

	// Pass the server, app, and pages to NewMenu
	menuInstance := menu.NewMenu(server, uiApp, menuPages)

	uiInstance := ui.NewUI()

	rootCmd := cmd.NewRootCmd(server, menuInstance, uiInstance)

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
