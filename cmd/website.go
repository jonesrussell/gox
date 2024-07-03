package cmd

import (
	"fmt"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
	"log"

	"github.com/rivo/tview"
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
			debug, err := rootCmd.Flags().GetBool("debug")
			if err != nil {
				log.Println("Can't get debug flag, defaulting to false")
				debug = false
			}

			if debug {
				log.Println("Debugging")
			}

			err = w.server.Start()
			if err != nil {
				log.Println("Error starting server:", err)
				return
			}

			uiApp := tview.NewApplication()
			uiPages := tview.NewPages()

			// Get the menu content as a tview.List.
			menuContent := w.menu.Display(uiApp, uiPages)

			// Create a TextView for the HTML representation of the website.
			htmlView := tview.NewTextView().SetText(w.server.GetHTML())

			// Add pages based on the menu options
			menuOptions := w.menu.GetOptions()
			for i, option := range menuOptions {
				func(i int, option string) {
					// Add a page with dummy content
					dummyContent := tview.NewTextView().SetText(fmt.Sprintf("This is page %s.", option))
					uiPages.AddPage(option, dummyContent, false, i == 0)
				}(i, option)
			}

			// Main flex layout
			flex := tview.NewFlex().
				// Left column (1/3 x width of screen)
				AddItem(menuContent, 0, 1, true).
				// Middle column (1/3 x width of screen)
				AddItem(uiPages, 0, 1, false).
				// Right column (1/3 x width of screen)
				AddItem(htmlView, 0, 1, false)

			if err := uiApp.SetRoot(flex, true).SetFocus(menuContent).Run(); err != nil {
				log.Fatalf("Error running application: %v", err)
			}
		},
	}

	return websiteCmd
}
