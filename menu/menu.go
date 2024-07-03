package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jonesrussell/gocreate/websiteserver"

	"github.com/rivo/tview"
)

type MenuInterface interface {
	Display(app *tview.Application, pages *tview.Pages) *tview.List
	handleChangeTitle(app *tview.Application, pages *tview.Pages)
	handleChangeBody(app *tview.Application, pages *tview.Pages)
	handleExit()
	GetOptions() []string
}

type menuImpl struct {
	reader  *bufio.Reader
	server  *websiteserver.WebsiteServerInterface
	options []string
}

// Ensure menuImpl implements MenuInterface
var _ MenuInterface = &menuImpl{}

func NewMenu(server *websiteserver.WebsiteServerInterface) *menuImpl {
	return &menuImpl{
		reader:  bufio.NewReader(os.Stdin),
		server:  server,
		options: []string{"Change title", "Update body", "Exit"},
	}
}

func (m *menuImpl) Display(app *tview.Application, pages *tview.Pages) *tview.List {
	// Create a new List.
	list := tview.NewList()

	// Use the List to display your menu.
	list.AddItem("Change title", "Press to change the title", '1', func() {
		m.handleChangeTitle(app, pages)
	}).
		AddItem("Update body", "Press to update the body", '2', func() {
			m.handleChangeBody(app, pages)
		}).
		AddItem("Exit", "Press to exit", '3', func() {
			m.handleExit()
		})

	return list
}

func (m *menuImpl) handleChangeTitle(app *tview.Application, pages *tview.Pages) {
	form := tview.NewForm()
	form.AddInputField("New title", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		// Get the text from the input field
		newTitle := form.GetFormItemByLabel("New title").(*tview.InputField).GetText()
		(*m.server).UpdateTitle(strings.TrimSpace(newTitle)) // Remove newline
		app.SetFocus(pages)                                  // Set focus back to the pages
	})
	form.SetBorder(true).SetTitle("Enter new title").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("ChangeTitle", form, true, true)
	app.SetFocus(form)
}

func (m *menuImpl) handleChangeBody(app *tview.Application, pages *tview.Pages) {
	form := tview.NewForm()
	form.AddInputField("New body", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		// Get the text from the input field
		newBody := form.GetFormItemByLabel("New body").(*tview.InputField).GetText()
		(*m.server).UpdateBody(strings.TrimSpace(newBody)) // Remove newline
		app.SetFocus(pages)                                // Set focus back to the pages
	})
	form.SetBorder(true).SetTitle("Enter new body").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("ChangeBody", form, true, true)
	app.SetFocus(form)
}

func (m *menuImpl) handleExit() {
	fmt.Println("Exiting...")
	(*m.server).Stop()
}

func (m *menuImpl) GetOptions() []string {
	return m.options
}
