package cmd

import (
	"fmt"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
	"log"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

func (w *WebsiteCommand) handleDebugFlag(flagset *pflag.FlagSet) bool {
	debug, err := flagset.GetBool("debug")
	if err != nil {
		log.Println("Can't get debug flag, defaulting to false")
		debug = false
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

func (w *WebsiteCommand) createFlexLayout(
	menuList *tview.List,
	menuPages *tview.Pages,
	htmlView *tview.TextView,
) *tview.Flex {
	return tview.NewFlex().
		// Left column (1/3 x width of screen)
		AddItem(menuList, 0, 1, true).
		// Middle column (1/3 x width of screen)
		AddItem(menuPages, 0, 1, false).
		// Right column (1/3 x width of screen)
		AddItem(htmlView, 0, 1, false)
}

func (w *WebsiteCommand) runApp(layout *tview.Flex) {
	app := w.menu.GetApp()
	fmt.Println("after call to w.menu.GetApp()")

	app.SetRoot(layout, true)
	fmt.Println("after call to app.SetRoot()")

	if err := app.Run(); err != nil {
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
			debug := w.handleDebugFlag(cmd.Flags())

			if debug {
				log.Println("Debugging")
			}

			if err := w.startServer(); err != nil {
				return
			}

			htmlView := tview.NewTextView().SetText(w.server.GetHTML())

			layout := w.createFlexLayout(
				w.menu.CreateMenu(),
				w.menu.GetPages(),
				htmlView,
			)

			w.runApp(layout)
		},
	}

	return websiteCmd
}
