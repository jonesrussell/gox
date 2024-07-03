package cmd

import (
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

func (w *WebsiteCommand) handleDebugFlag(cmd *cobra.Command) bool {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Println("Can't get debug flag, defaulting to false")
		debug = false
	}

	if debug {
		log.Println("Debugging")
	}

	return debug
}

func (w *WebsiteCommand) startServer() error {
	err := w.server.Start()
	if err != nil {
		log.Println("Error starting server:", err)
		return err
	}
	return nil
}

func (w *WebsiteCommand) createMenuList() *tview.List {
	return w.menu.CreateMenu()
}

func (w *WebsiteCommand) createHTMLView() *tview.TextView {
	return tview.NewTextView().SetText(w.server.GetHTML())
}

func (w *WebsiteCommand) createFlexLayout(menuList *tview.List, menuPages *tview.Pages, htmlView *tview.TextView) *tview.Flex {
	return tview.NewFlex().
		// Left column (1/3 x width of screen)
		AddItem(menuList, 0, 1, true).
		// Middle column (1/3 x width of screen)
		AddItem(menuPages, 0, 1, false).
		// Right column (1/3 x width of screen)
		AddItem(htmlView, 0, 1, false)
}

func (w *WebsiteCommand) runApp(layout *tview.Flex, menuList *tview.List) {
	if err := w.menu.GetApp().SetRoot(layout, true).SetFocus(menuList).Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
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
			w.handleDebugFlag(cmd)

			if err := w.startServer(); err != nil {
				return
			}

			menuList := w.createMenuList()
			menuPages := w.menu.GetPages()
			htmlView := w.createHTMLView()

			layout := w.createFlexLayout(
				menuList,
				menuPages,
				htmlView,
			)

			w.runApp(layout, menuList)
		},
	}

	return websiteCmd
}
