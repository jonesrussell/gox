package websiteserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// Create a new mock server
	server := NewMockServer(&Page{})

	// Assert that the server was created
	assert.NotNil(t, server)
}
