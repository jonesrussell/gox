package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/debug"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
	"log"

	"github.com/rivo/tview"
)

func main() {
	// Create a new LogDebugger
	debugger := debug.NewLogDebugger()
	err := debugger.Init()
	if err != nil {
		log.Fatalf("Error initializing debugger: %v", err)
	}

	// Pass the Debugger to NewServer
	server := websiteserver.NewServer(debugger)

	// Create tview.Application and tview.Pages instances
	uiApp := tview.NewApplication()
	menuPages := tview.NewPages()

	// Pass the server, app, and pages to NewMenu
	menuInstance := menu.NewMenu(server, uiApp, menuPages)

	rootCmd := cmd.NewRootCmd(server, menuInstance)

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
