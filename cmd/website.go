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
				title := promptForTitle()
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

func promptForTitle() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter new title: ")
	newTitle, _ := reader.ReadString('\n')
	return strings.TrimSpace(newTitle) // Remove newline
}

func init() {
	rootCmd.AddCommand(websiteCmd)
}
