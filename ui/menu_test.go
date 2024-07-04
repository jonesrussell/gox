package ui

import (
	"testing"

	"jonesrussell/gocreate/webserver"

	"github.com/stretchr/testify/assert"
)

func TestNewMenu(t *testing.T) {
	// Mock the server
	server := webserver.NewMockServer(&webserver.Page{})

	// Create a new menu
	menuInstance := NewMenu(server, nil, nil)

	// Assert that the menu was created with the correct options
	assert.Equal(t, []string{"Change title", "Update body", "Exit"}, menuInstance.GetOptions())
}

func TestCreateMenu(t *testing.T) {
	// Mock the server
	server := webserver.NewMockServer(&webserver.Page{})

	// Create a new menu
	menuInstance := NewMenu(server, nil, nil)

	// Create the menu
	list := menuInstance.CreateMenu()

	// Assert that the list has the correct number of items
	assert.Equal(t, 3, list.GetItemCount())
}
