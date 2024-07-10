package cmd

import (
	"fmt"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"log"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/yosssi/gohtml"
)

type WebsiteCommand struct {
	server     webserver.WebServerInterface
	menu       ui.MenuInterface
	ui         ui.UIInterface
	updateChan <-chan struct{}
	htmlView   *tview.TextView
}

func NewWebsiteCommand(
	server webserver.WebServerInterface,
	menu ui.MenuInterface,
	ui ui.UIInterface,
) *WebsiteCommand {
	htmlView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			ui.GetApp().Draw()
		})

	cmd := &WebsiteCommand{
		server:     server,
		menu:       menu,
		ui:         ui,
		updateChan: server.GetUpdateChan(),
		htmlView:   htmlView,
	}

	// Initialize HTML content
	cmd.updateHTMLView()

	return cmd
}

func (w *WebsiteCommand) HandleDebugFlag(flagset *pflag.FlagSet) bool {
	debug, err := flagset.GetBool("debug")
	if err != nil {
		log.Println("Can't get debug flag, defaulting to false")
		debug = false
	}

	return debug
}

func (w *WebsiteCommand) StartServer() error {
	err := w.server.Start()
	if err != nil {
		w.server.Logger().Debug("Error starting server")
		log.Fatalf("Error starting server: %v", err)
	} else {
		w.server.Logger().Debug("Server started successfully")
	}
	return nil
}

func (w *WebsiteCommand) ConfigureAddressView(addressView *tview.TextView) {
	// Add a border and title to the addressView
	addressView.SetBorder(true).
		SetTitle("Webserver")
}

func (w *WebsiteCommand) CreateThirdColumn(
	htmlView *tview.TextView,
	addressView *tview.TextView,
) *tview.Flex {
	// Create a new Flex for the third column
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		// addressView takes up 1/7th of the height
		AddItem(addressView, 0, 1, false).
		// htmlView takes up the rest of the height (7/8th)
		AddItem(htmlView, 0, 6, false)
}

func (w *WebsiteCommand) CreateFlexLayout(
	server webserver.WebServerInterface,
	uiApp *tview.Application,
	menuPages *tview.Pages,
	htmlView *tview.TextView,
	addressView *tview.TextView,
) *tview.Flex {
	w.ConfigureAddressView(addressView)
	thirdColumn := w.CreateThirdColumn(htmlView, addressView)

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
		log.Fatalf("Error creating UI: %v", err)
	}
}

func (w *WebsiteCommand) updateHTMLView() {
	html := w.server.GetHTML()
	formattedHTML := gohtml.Format(html)
	w.htmlView.Clear()
	fmt.Fprintf(w.htmlView, "%s", formattedHTML)
}

func (w *WebsiteCommand) startHTMLViewUpdater() {
	go func() {
		for range w.updateChan {
			w.ui.GetApp().QueueUpdateDraw(func() {
				w.updateHTMLView()
			})
		}
	}()
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
			debug := w.HandleDebugFlag(cmd.Flags())

			if debug {
				log.Println("Debugging")
			}

			if err := w.StartServer(); err != nil {
				return
			}

			// Start the HTML view updater
			w.startHTMLViewUpdater()

			// Create a TextView for the server address
			addressView := tview.NewTextView().SetText(w.server.GetURL())

			layout := w.CreateFlexLayout(
				w.server,
				w.ui.GetApp(),
				w.ui.GetPages(),
				w.htmlView, // Use w.htmlView instead of creating a new one
				addressView,
			)

			w.runApp(layout)
		},
	}

	return websiteCmd
}
