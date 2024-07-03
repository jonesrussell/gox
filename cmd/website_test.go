package cmd_test

import (
	"fmt"
	"testing"

	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/debug"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/utils"
	"jonesrussell/gocreate/websiteserver"

	"github.com/stretchr/testify/assert"
)

// Create a new LogDebugger and WebsiteUpdater once for all tests
var (
	debugger = debug.NewLogDebugger()
	updater  = websiteserver.NewWebsiteUpdater(debugger)
)

func Test_ServerInitialization(t *testing.T) {
	mockPage := websiteserver.NewPage("", "", utils.MockFileReader{}, updater)
	mockServer := websiteserver.NewMockServer(mockPage)

	_, ok := mockServer.(*websiteserver.MockServer)
	assert.True(t, ok, "Expected a mock server, got a real one")
}

func Test_WebsiteCommand(t *testing.T) {
	mockPage := websiteserver.NewPage("", "", utils.MockFileReader{}, updater)
	server := websiteserver.NewMockServer(mockPage)
	m := menu.NewMockMenu(&server)

	websiteCommand := cmd.NewWebsiteCommand(server, m)
	websiteCmd := websiteCommand.Command()

	output, err := cmd.ExecuteCommand(websiteCmd, "website")

	fmt.Println(output)
	foo := debug.NewLogDebugger()
	foo.Debug(output)

	assert.NoError(t, err, "Expected no error")
	assert.Empty(t, output, "Expected output to match")

	// mockMenu, ok := m.(*menu.MockMenu)
	// assert.True(t, ok, "Expected m to be of type *MockMenu")
	// assert.True(t, mockMenu.ExitCalled, "Expected handleExit to be called")
}
