package cmd_test

import (
	"testing"

	"jonesrussell/gocreate/debug"
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
