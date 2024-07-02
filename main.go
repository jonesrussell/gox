package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/debug"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
)

func main() {
	// Create a new LogDebugger
	debugger := &debug.LogDebugger{}
	debugger.Init()

	// Pass the Debugger to NewServer
	server := websiteserver.NewServer(debugger)

	menuInstance := menu.NewMenu(&server)
	rootCmd := cmd.NewRootCmd(server, menuInstance)
	rootCmd.Execute()
}
