package cmd_test

import (
	"fmt"
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/utils"
	"jonesrussell/gocreate/websiteserver"
	"testing"

	"github.com/spf13/cobra"
)

func TestWebsiteCmdAddedOnce(t *testing.T) {
	fmt.Println("Starting TestWebsiteCmdAddedOnce")

	rootCmd := &cobra.Command{
		Use: "root",
	}

	mockPage := websiteserver.NewPage("", utils.MockFileReader{})
	mockServer := websiteserver.NewMockServer(mockPage)
	m := menu.NewMockMenu(&mockServer)

	// Manually add the websiteCmd to rootCmd
	rootCmd.AddCommand(cmd.NewWebsiteCommand(mockServer, m.(*menu.Menu)))

	var count int
	for _, c := range rootCmd.Commands() {
		if c.Name() == "website" {
			count++
		}
	}

	fmt.Printf("Found 'website' command %d times\n", count)

	if count != 1 {
		t.Errorf("Expected 'website' command to be added once to rootCmd, but got %d", count)
	}

	fmt.Println("Finished TestWebsiteCmdAddedOnce")
}
