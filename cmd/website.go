// cmd/website.go
package cmd

import (
	"fmt"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"

	"github.com/spf13/cobra"
)

type WebsiteCommand struct {
	server websiteserver.WebsiteServerInterface
	menu   menu.MenuInterface
}

func NewWebsiteCommand(
	server websiteserver.WebsiteServerInterface,
	menu menu.MenuInterface,
) *WebsiteCommand {
	return &WebsiteCommand{
		server: server,
		menu:   menu,
	}
}

func (w *WebsiteCommand) Command() *cobra.Command {
	websiteCmd := &cobra.Command{
		Use:   "website",
		Short: "Create a no code website from the command line",
		Long: `The 'website' command allows you to create a
			no-code website directly from the command line. This is
			particularly useful for beginner developers who need to
			quickly	set up a static website.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := w.server.Start()
			if err != nil {
				fmt.Println("Error starting server:", err)
				return
			}

			w.menu.Display()
		},
	}

	return websiteCmd
}
