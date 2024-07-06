package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"testing"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

var (
	mockServer = webserver.NewMockServer(nil)
	mockMenu   = ui.NewMockMenu(mockServer)
	mockUI     = ui.NewMockUI()
)

func TestNewWebsiteCommand(t *testing.T) {
	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)

	assert.NotNil(t, websiteCmd)
	assert.IsType(t, &cmd.WebsiteCommand{}, websiteCmd)
}

func TestWebsiteCommand_Command(t *testing.T) {
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)
	command := websiteCmd.Command()

	assert.NotNil(t, command)
	assert.IsType(t, &cobra.Command{}, command)
	assert.Equal(t, "website", command.Use)
	assert.Contains(t, command.Short, "Create a no code website")
}

func TestWebsiteCommand_HandleDebugFlag(t *testing.T) {
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)

	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.Bool("debug", false, "debug flag")

	// Test when debug flag is not set
	debug := websiteCmd.HandleDebugFlag(flagSet)
	assert.False(t, debug)

	// Test when debug flag is set
	flagSet.Set("debug", "true")
	debug = websiteCmd.HandleDebugFlag(flagSet)
	assert.True(t, debug)
}

// func TestWebsiteCommand_StartServer(t *testing.T) {
// 	mockLogger := new(logger.MockLogger)
// 	mockServer := webserver.NewMockServer(nil)
// 	mockMenu := ui.NewMockMenu(&mockServer)
// 	mockUI := ui.NewMockUI()

// 	// Set expectations
// 	mockLogger.On("Debug", mock.Anything).Return()

// 	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)

// 	err := websiteCmd.StartServer()
// 	assert.NoError(t, err)

// 	mockLogger.AssertExpectations(t)
// }

func TestWebsiteCommand_ConfigureAddressView(t *testing.T) {
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)

	addressView := tview.NewTextView()
	websiteCmd.ConfigureAddressView(addressView)

	// Check if the border is set
	// TextView doesn't have a GetBorder() method, so we can't directly check this
	// Instead, we can check if the border color is set, which implies a border
	assert.NotEqual(t, tview.Styles.PrimitiveBackgroundColor, addressView.GetBorderColor())

	// Check the title
	assert.Equal(t, "Webserver", addressView.GetTitle())

	// Optionally, check other properties that you set in ConfigureAddressView
	// For example, if you set text alignment:
	// assert.Equal(t, tview.AlignCenter, addressView.GetTextAlign())
}

func TestWebsiteCommand_CreateThirdColumn(t *testing.T) {
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)

	htmlView := tview.NewTextView()
	addressView := tview.NewTextView()

	thirdColumn := websiteCmd.CreateThirdColumn(htmlView, addressView)

	assert.NotNil(t, thirdColumn)
	assert.IsType(t, &tview.Flex{}, thirdColumn)
	assert.Equal(t, 2, thirdColumn.GetItemCount())
}

func TestWebsiteCommand_createFlexLayout(t *testing.T) {
	mockServer := webserver.NewMockServer(nil)
	mockMenu := ui.NewMockMenu(mockServer)
	mockUI := ui.NewMockUI()

	websiteCmd := cmd.NewWebsiteCommand(mockServer, mockMenu, mockUI)

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
