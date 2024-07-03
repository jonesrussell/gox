package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/debug"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/utils"
	"jonesrussell/gocreate/websiteserver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WebsiteCommand(t *testing.T) {
	// Create a new LogDebugger
	debugger := &debug.LogDebugger{}
	debugger.Init()

	// Create a new WebsiteUpdater
	updater := websiteserver.NewWebsiteUpdater(debugger)

	mockPage := websiteserver.NewPage("", "", utils.MockFileReader{}, updater)
	server := websiteserver.NewMockServer(mockPage)
	m := menu.NewMockMenu(&server)

	// Create the WebsiteCommand with the mocked dependencies
	websiteCommand := cmd.NewWebsiteCommand(server, m)

	// Get the *cobra.Command from the WebsiteCommand
	websiteCmd := websiteCommand.Command()

	// Execute the command with the desired arguments
	// Assuming ExecuteCommand is a function that executes the command and returns output and error
	output, err := cmd.ExecuteCommand(websiteCmd, "website")

	// Assert something about the output or error
	assert.NoError(t, err, "Expected no error")

	// Assert something about the output
	expectedOutput := ""
	assert.Equal(t, expectedOutput, output, "Expected output to match")

	// Assert that handleExit was called
	mockMenu, ok := m.(*menu.MockMenu)
	if !ok {
		t.Fatal("Expected m to be of type *MockMenu")
	}
	assert.True(t, mockMenu.ExitCalled, "Expected handleExit to be called")
}
