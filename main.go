package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/debug"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"

	"github.com/rivo/tview"
)

func main() {
	// Create a new LogDebugger
	debugger := &debug.LogDebugger{}
	debugger.Init()

	// Pass the Debugger to NewServer
	server := websiteserver.NewServer(debugger)

	// Create tview.Application and tview.Pages instances
	uiApp := tview.NewApplication()
	menuPages := tview.NewPages()

	// Pass the server, app, and pages to NewMenu
	menuInstance := menu.NewMenu(server, uiApp, menuPages)

	rootCmd := cmd.NewRootCmd(server, menuInstance)
	rootCmd.Execute()
}
