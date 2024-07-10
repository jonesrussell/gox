package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"testing"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestWebsiteCommand_Command(t *testing.T) {
	cfg := &cmd.Config{
		Debug: false,
	}
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(cfg, mockServer, mockMenu, mockUI)
	command := websiteCmd.Command()

	assert.NotNil(t, command)
	assert.IsType(t, &cobra.Command{}, command)
	assert.Equal(t, "website", command.Use)
	assert.Contains(t, command.Short, "Create a no code website")
}

func TestWebsiteCommand_HandleDebugFlag(t *testing.T) {
	cfg := &cmd.Config{
		Debug: true,
	}
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(cfg, mockServer, mockMenu, mockUI)

	assert.True(t, websiteCmd.Config.Debug)
}

func TestWebsiteCommand_ConfigureAddressView(t *testing.T) {
	cfg := &cmd.Config{
		Debug: true,
	}
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(cfg, mockServer, mockMenu, mockUI)

	addressView := tview.NewTextView()
	websiteCmd.ConfigureAddressView(addressView)

	assert.NotEqual(t, tview.Styles.PrimitiveBackgroundColor, addressView.GetBorderColor())
	assert.Equal(t, "Webserver", addressView.GetTitle())
}

func TestWebsiteCommand_CreateThirdColumn(t *testing.T) {
	cfg := &cmd.Config{
		Debug: true,
	}
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(cfg, mockServer, mockMenu, mockUI)

	htmlView := tview.NewTextView()
	addressView := tview.NewTextView()

	thirdColumn := websiteCmd.CreateThirdColumn(htmlView, addressView)

	assert.NotNil(t, thirdColumn)
	assert.IsType(t, &tview.Flex{}, thirdColumn)
	assert.Equal(t, 2, thirdColumn.GetItemCount())
}

func TestWebsiteCommand_createFlexLayout(t *testing.T) {
	cfg := &cmd.Config{
		Debug: true,
	}
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(cfg, mockServer, mockMenu, mockUI)

	app := tview.NewApplication()
	pages := tview.NewPages()
	htmlView := tview.NewTextView()
	addressView := tview.NewTextView()

	mockUI.SetApp(app)
	mockUI.SetPages(pages)

	layout := websiteCmd.CreateFlexLayout(mockServer, app, pages, htmlView, addressView)

	assert.NotNil(t, layout)
	assert.IsType(t, &tview.Flex{}, layout)
	assert.Equal(t, 3, layout.GetItemCount())
}
