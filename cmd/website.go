package cmd

import (
	"jonesrussell/gocreate/ui" // Import the ui package
	"jonesrussell/gocreate/webserver"
	"log"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type WebsiteCommand struct {
	server webserver.WebServerInterface
	menu   ui.MenuInterface
	ui     ui.UIInterface
}

func NewWebsiteCommand(
	server webserver.WebServerInterface,
	menu ui.MenuInterface,
	ui ui.UIInterface,
) *WebsiteCommand {
	return &WebsiteCommand{
		server: server,
		menu:   menu,
		ui:     ui,
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
		w.server.Logger().Debug("Error starting server")
		log.Fatalf("Error starting server: %v", err)
	} else {
		w.server.Logger().Debug("Server started successfully")
	}
	return nil
}

func (w *WebsiteCommand) configureAddressView(addressView *tview.TextView) {
	// Add a border and title to the addressView
	addressView.SetBorder(true).SetTitle("Webserver")
}

func (w *WebsiteCommand) createThirdColumn(
	htmlView *tview.TextView,
	addressView *tview.TextView,
) *tview.Flex {
	// Create a new Flex for the third column
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		// addressView takes up 1/8th of the height
		AddItem(addressView, 0, 1, false).
		// htmlView takes up the rest of the height (7/8th)
		AddItem(htmlView, 0, 6, false)
}

func (w *WebsiteCommand) createFlexLayout(
	server webserver.WebServerInterface,
	uiApp *tview.Application,
	menuPages *tview.Pages,
	htmlView *tview.TextView,
	addressView *tview.TextView,
) *tview.Flex {
	w.configureAddressView(addressView)
	thirdColumn := w.createThirdColumn(htmlView, addressView)

	// Use the Menu struct's CreateMenu() method to create and populate the menu
	menu := ui.NewMenu(server, uiApp, menuPages)
	menuList := menu.CreateMenu()

	return tview.NewFlex().
		// Left column (1/3 x width of screen)
		AddItem(menuList, 0, 1, true).
		// Middle column (1/3 x width of screen)
		AddItem(menuPages, 0, 1, false).
		// Right column (1/3 x width of screen)
		AddItem(thirdColumn, 0, 1, false)
}

func (w *WebsiteCommand) runApp(layout *tview.Flex) {
	if err := w.ui.GetApp().SetRoot(layout, true).Run(); err != nil {
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

			// Create a TextView for the server address
			addressView := tview.NewTextView().SetText(w.server.GetURL())

			layout := w.createFlexLayout(
				w.server,
				w.ui.GetApp(),
				w.ui.GetPages(),
				htmlView,
				addressView,
			)

			w.runApp(layout)
		},
	}

	return websiteCmd
}
