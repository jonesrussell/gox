package ui

import (
	"github.com/rivo/tview"
)

// UIInterface is the interface for UI operations
type UIInterface interface {
	Run() error
}

// UI implements the UIInterface
type UI struct {
	// Add fields as needed
	app *tview.Application
}

// Ensure UI implements UIInterface
var _ UIInterface = &UI{}

// NewUI creates a new UI instance
func NewUI() *UI {
	return &UI{
		app: tview.NewApplication(),
	}
}

// Run starts the UI application
func (ui *UI) Run() error {
	// Implement your UI logic here
	return nil
}
