package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
	"testing"

	"github.com/spf13/cobra"
)

func TestWebsiteCmdAddedOnce(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "root",
	}

	mockServer := websiteserver.NewMockServer()
	m := menu.NewMockMenu(&mockServer)

	// Manually add the websiteCmd to rootCmd
	rootCmd.AddCommand(cmd.NewWebsiteCommand(mockServer, m.(*menu.Menu)))

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
