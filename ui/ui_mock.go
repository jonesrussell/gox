package ui

import (
	"github.com/rivo/tview"
)

// MockUI is a mock implementation of UIInterface for testing
type MockUI struct {
	app   *tview.Application
	menu  *tview.List
	pages *tview.Pages
}

// Ensure MockUI implements UIInterface
var _ UIInterface = &MockUI{}

// NewMockUI creates a new MockUI instance
func NewMockUI() *MockUI {
	return &MockUI{
		app:   tview.NewApplication(),
		menu:  tview.NewList(),
		pages: tview.NewPages(),
	}
}

func (m *MockUI) CreateMenu() *tview.List {
	return m.menu
}

func (m *MockUI) GetPages() *tview.Pages {
	return m.pages
}

func (m *MockUI) GetApp() *tview.Application {
	return m.app
}

// Run mocks starting the UI application
func (m *MockUI) Run() error {
	// Mock implementation
	return nil
}

// Additional methods for setting up expectations or verifying calls can be added here
func (m *MockUI) SetMenu(menu *tview.List) {
	m.menu = menu
}

func (m *MockUI) SetPages(pages *tview.Pages) {
	m.pages = pages
}

func (m *MockUI) SetApp(app *tview.Application) {
	m.app = app
}
