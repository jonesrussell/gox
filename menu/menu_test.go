package menu_test

import (
	"testing"

	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"

	"github.com/stretchr/testify/assert"
)

func TestCreateMenu(t *testing.T) {
	mockServer := websiteserver.NewMockServer()
	mockMenu := menu.NewMockMenu(&mockServer)

	list := mockMenu.CreateMenu()

	assert.NotNil(t, list, "Expected list to be not nil")
}

func TestHandleExit(t *testing.T) {
	mockServer := websiteserver.NewMockServer()
	mockMenu := menu.NewMockMenu(&mockServer)

	mockMenu.HandleExit()

	assert.True(t, mockMenu.ExitCalled, "Expected ExitCalled to be true")
}

func TestGetOptions(t *testing.T) {
	mockServer := websiteserver.NewMockServer()
	mockMenu := menu.NewMockMenu(&mockServer)

	options := mockMenu.GetOptions()

	expectedOptions := []string{"Option1", "Option2"}
	assert.Equal(t, expectedOptions, options, "Expected options to be equal to expectedOptions")
}

func TestGetApp(t *testing.T) {
	mockServer := websiteserver.NewMockServer()
	mockMenu := menu.NewMockMenu(&mockServer)

	app := mockMenu.GetApp()

	assert.NotNil(t, app, "Expected app to be not nil")
}

func TestGetPages(t *testing.T) {
	mockServer := websiteserver.NewMockServer()
	mockMenu := menu.NewMockMenu(&mockServer)

	pages := mockMenu.GetPages()

	assert.NotNil(t, pages, "Expected pages to be not nil")
}
