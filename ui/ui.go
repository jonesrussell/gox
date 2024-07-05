package ui

import (
	"github.com/rivo/tview"
)

// UIInterface is the interface for UI operations
type UIInterface interface {
	CreateMenu() *tview.List
	GetPages() *tview.Pages
	GetApp() *tview.Application
	Run() error
}

// UI implements the UIInterface
type UI struct {
	app   *tview.Application
	menu  *tview.List
	pages *tview.Pages
}

// Ensure UI implements UIInterface
var _ UIInterface = &UI{}

// NewUI creates a new UI instance
func NewUI() *UI {
	return &UI{
		app:   tview.NewApplication(),
		menu:  tview.NewList(),
		pages: tview.NewPages(),
	}
}

func (u *UI) CreateMenu() *tview.List {
	// Implement the CreateMenu method
	return u.menu
}

func (u *UI) GetPages() *tview.Pages {
	// Implement the GetPages method
	return u.pages
}
func (u *UI) GetApp() *tview.Application {
	// Implement the GetApp method
	return u.app
}

// Run starts the UI application
func (ui *UI) Run() error {
	// Implement your UI logic here
	return nil
}
