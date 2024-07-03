package menu

import (
	"bufio"
	"fmt"
	"jonesrussell/gocreate/websiteserver"
	"log"

	"github.com/rivo/tview"
)

// MockMenu simulates the behavior of the real Menu for testing purposes.
type MockMenu struct {
	reader *bufio.Reader
	server *websiteserver.WebsiteServerInterface
	// Add fields to capture method calls if needed
	ExitCalled bool
}

// Ensure MockMenu implements MenuInterface
var _ MenuInterface = &MockMenu{}

// NewMockMenu creates a new instance of MockMenu.
func NewMockMenu(server *websiteserver.WebsiteServerInterface) MenuInterface {
	log.Println("NewMockMenu method called")

	return &MockMenu{
		reader: bufio.NewReader(nil), // Control the input
		server: server,
	}
}

// Display mocks the interaction with the user.
func (m *MockMenu) Display(app *tview.Application, pages *tview.Pages) *tview.List {
	log.Println("Display method called")
	list := tview.NewList()
	log.Println("after tview.NewList() call")
	// fmt.Fprintln(tv, "Mock menu displayed.")
	return list
}

// handleChangeTitle mocks changing the title.
func (m *MockMenu) handleChangeTitle(app *tview.Application, pages *tview.Pages) {
	// Implementation without returning an error
}

func (m *MockMenu) handleChangeBody(app *tview.Application, pages *tview.Pages) {
	// Your implementation here
}

// handleExit mocks exiting the menu.
func (m *MockMenu) handleExit() {
	fmt.Println("Exiting...")
	m.ExitCalled = true
}

func (m *MockMenu) GetOptions() []string {
	// Implement the method according to your needs
	return []string{"Option1", "Option2"} // Example return value
}
