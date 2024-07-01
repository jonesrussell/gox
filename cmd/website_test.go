package cmd_test

import (
	"jonesrussell/gocreate/cmd"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/utils"
	"jonesrussell/gocreate/websiteserver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WebsiteCommand(t *testing.T) {
	// Assuming NewMockServer and NewMockMenu are correctly implemented elsewhere
	mockPage := websiteserver.NewPage("", utils.MockFileReader{})
	server := websiteserver.NewMockServer(mockPage)
	m := menu.NewMockMenu(&server)

	// Create the command with the mocked dependencies
	websiteCmd := cmd.NewWebsiteCommand(server, m)

	// Execute the command with the desired arguments
	// Assuming ExecuteCommand is a function that executes the command and returns output and error
	output, err := cmd.ExecuteCommand(websiteCmd, "website")

	// Assert something about the output or error
	assert.NoError(t, err, "Expected no error")

	// Assert something about the output
	expectedOutput := "Expected output"
	assert.Equal(t, expectedOutput, output, "Expected output to match")
}
