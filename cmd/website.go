package cmd

import (
	"bufio"
	"fmt"
	"jonesrussell/gocreate/httpserver"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var websiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Create a no code website from the command line",
	Long: `The 'website' command allows you to create a
	no-code website directly from the command line. This is
	particularly useful for beginner developers who need to
	quickly	set up a static website.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		server := httpserver.NewServer()

		server.Start()

		for {
			fmt.Println("\nInteractive Menu:")
			fmt.Println("1. Change title")
			fmt.Println("2. Exit")
			fmt.Print("Enter command number: ")

			command, _ := reader.ReadString('\n')
			command = strings.TrimSpace(command) // Remove newline

			switch command {
			case "1":
				title, err := promptForTitle()
				if err != nil {
					fmt.Println("Error updating title:", err)
					continue
				}
				server.UpdateTitle(title)
			case "2":
				fmt.Println("Exiting...")
				server.Stop()
				return
			default:
				fmt.Println("Invalid command. Please enter a number from 1 to 2.")
			}
		}
	},
}

func promptForTitle() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter new title: ")
	newTitle, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(newTitle), nil // Remove newline
}

func init() {
	rootCmd.AddCommand(websiteCmd)
}
