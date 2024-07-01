package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
)

func main() {
	server := websiteserver.NewServer()             // Replace with actual instantiation logic
	menuInstance := menu.NewMenu(&server)           // Replace with actual instantiation logic
	rootCmd := cmd.NewRootCmd(server, menuInstance) // Replace with actual instantiation logic
	rootCmd.Execute()
}
