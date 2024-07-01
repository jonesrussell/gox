// cmd/website.go
package cmd

import (
	"fmt"
	"jonesrussell/gocreate/httpserver"
	"jonesrussell/gocreate/menu"

	"github.com/spf13/cobra"
)

var WebsiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Create a no code website from the command line",
	Long: `The 'website' command allows you to create a
	no-code website directly from the command line. This is
	particularly useful for beginner developers who need to
	quickly	set up a static website.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := httpserver.NewServer()

		err := server.Start()
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}

		m := menu.NewMenu(server)
		m.Display()
	},
}

func init() {
	rootCmd.AddCommand(WebsiteCmd)
}
