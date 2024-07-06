package cmd_test

import (
	"testing"

	"jonesrussell/gocreate/logger"
	"jonesrussell/gocreate/utils"
	"jonesrussell/gocreate/webserver"

	"github.com/stretchr/testify/assert"
)

// Create a new Logger and WebsiteUpdater once for all tests
var (
	weblogger, _ = logger.NewLogger("/tmp/gocreate-tests.log")
	updater      = webserver.NewWebsiteUpdater(weblogger)
)

func Test_ServerInitialization(t *testing.T) {
	mockPage := webserver.NewPage("", "", utils.MockFileReader{}, updater, "static/index.html")
	mockServer := webserver.NewMockServer(mockPage)

	_, ok := mockServer.(*webserver.MockServer)
	assert.True(t, ok, "Expected a mock server, got a real one")
}
