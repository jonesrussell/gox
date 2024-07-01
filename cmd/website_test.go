package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"testing"

	"github.com/spf13/cobra"
)

func TestWebsiteCmdAddedOnce(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "root",
	}

	// Manually add the websiteCmd to rootCmd
	rootCmd.AddCommand(cmd.WebsiteCmd)

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

func TestWebsiteCmdNotAddedWithDifferentName(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "root",
	}

	differentCmd := &cobra.Command{
		Use: "different",
	}

	rootCmd.AddCommand(differentCmd)

	// Manually add the websiteCmd to rootCmd
	rootCmd.AddCommand(cmd.WebsiteCmd)

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

func TestWebsiteCmdNotAddedWithNilRootCmd(t *testing.T) {
	var rootCmd *cobra.Command

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when passing nil rootCmd to AddCommand function")
		}
	}()

	rootCmd.AddCommand(cmd.WebsiteCmd)
}
