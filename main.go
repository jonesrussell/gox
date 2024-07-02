package main

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
)

func main() {
	server := websiteserver.NewServer()
	menuInstance := menu.NewMenu(&server)
	rootCmd := cmd.NewRootCmd(server, menuInstance)
	rootCmd.Execute()
}
