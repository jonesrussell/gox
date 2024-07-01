package menu

import (
	"bufio"
	"jonesrussell/gocreate/websiteserver"
)

// MockMenu simulates the behavior of the real Menu for testing purposes.
type MockMenu struct {
	reader *bufio.Reader
	server *websiteserver.WebsiteServerInterface
	// Add fields to capture method calls if needed
}

// Ensure MockMenu implements MenuInterface
var _ MenuInterface = &MockMenu{}

// NewMockMenu creates a new instance of MockMenu.
func NewMockMenu(server *websiteserver.WebsiteServerInterface) MenuInterface {
	return &MockMenu{
		reader: bufio.NewReader(nil), // Control the input
		server: server,
	}
}

// Display mocks the interaction with the user.
func (m *MockMenu) Display() {
	// Implement mocked behavior
}

// handleChangeTitle mocks changing the title.
func (m *MockMenu) handleChangeTitle() error {
	// Return a predefined title or simulate user input
	return nil
}

// handleExit mocks exiting the menu.
func (m *MockMenu) handleExit() {
	// Implement exit logic
}
