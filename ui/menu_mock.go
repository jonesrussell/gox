package ui

import (
	"bufio"
	"jonesrussell/gocreate/webserver"

	"github.com/rivo/tview"
)

// MockMenu simulates the behavior of the real Menu for testing purposes.
type MockMenu struct {
	reader *bufio.Reader
	server *webserver.WebServerInterface
	// Add fields to capture method calls if needed
	ExitCalled bool
}

// Ensure MockMenu implements MenuInterface
var _ MenuInterface = &MockMenu{}

// NewMockMenu creates a new instance of MockMenu.
func NewMockMenu(server *webserver.WebServerInterface) MenuInterface {
	return &MockMenu{
		reader: bufio.NewReader(nil), // Control the input
		server: server,
	}
}

// Display mocks the interaction with the user.
func (m *MockMenu) CreateMenu() *tview.List {
	list := tview.NewList()
	return list
}

// handleChangeTitle mocks changing the title.
func (m *MockMenu) handleChangeTitle() {
	// Implementation without returning an error
}

func (m *MockMenu) handleChangeBody() {
	// Your implementation here
}

// handleExit mocks exiting the ui.
func (m *MockMenu) handleExit() {
	m.ExitCalled = true
}

func (m *MockMenu) GetOptions() []string {
	// Implement the method according to your needs
	return []string{"Option1", "Option2"} // Example return value
}

func (m *MockMenu) GetApp() *tview.Application {
	// Return a mock or actual Application object
	app := tview.NewApplication() // Example return value
	return app
}

func (m *MockMenu) GetPages() *tview.Pages {
	// Create a new Pages object
	pages := tview.NewPages()

	// Populate the pages as needed for your test/mock scenario

	return pages
}
