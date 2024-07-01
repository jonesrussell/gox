package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/httpserver"
	"jonesrussell/gocreate/menu"
	"testing"

	"github.com/spf13/cobra"
)

func TestWebsiteCmdAddedOnce(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "root",
	}

	server := httpserver.NewServer()
	m := menu.NewMenu(&server)

	// Manually add the websiteCmd to rootCmd
	rootCmd.AddCommand(cmd.NewWebsiteCommand(server, m))

	var count int
	for _, c := range rootCmd.Commands() {
		if c.Name() == "website" {
			count++
		}
	}

	if count != 1 {
		t.Errorf("Expected 'website' command to be added once to rootCmd, but got %d", count)
	}
}
