// cmd/website.go
package cmd

import (
	"fmt"
	"jonesrussell/gocreate/httpserver"
	"jonesrussell/gocreate/menu"

	"github.com/spf13/cobra"
)

type WebsiteCommand struct {
	server httpserver.Server
	menu   *menu.Menu
}

func NewWebsiteCommand(server httpserver.Server, menu *menu.Menu) *cobra.Command {
	wc := &WebsiteCommand{server: server, menu: menu}

	return &cobra.Command{
		Use:   "website",
		Short: "Create a no code website from the command line",
		Long: `The 'website' command allows you to create a
			no-code website directly from the command line. This is
			particularly useful for beginner developers who need to
			quickly	set up a static website.`,
		Run: wc.run,
	}
}

func (wc *WebsiteCommand) run(cmd *cobra.Command, args []string) {
	err := wc.server.Start()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	wc.menu.Display()
}

func init() {
	server := httpserver.NewServer()
	m := menu.NewMenu(&server)
	rootCmd.AddCommand(NewWebsiteCommand(server, m))
}
