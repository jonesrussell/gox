package ui

import (
	"bufio"
	"log"
	"os"
	"strings"

	"jonesrussell/gocreate/webserver"

	"github.com/rivo/tview"
)

type MenuInterface interface {
	CreateMenu() *tview.List
	handleChangeTitle()
	handleChangeBody()
	handleExit()
	GetOptions() []string
	GetApp() *tview.Application
	GetPages() *tview.Pages
}

type Menu struct {
	reader    *bufio.Reader
	server    webserver.WebServerInterface
	options   []string
	uiApp     *tview.Application
	menuPages *tview.Pages
}

// Ensure Menu implements MenuInterface
var _ MenuInterface = &Menu{}

func NewMenu(
	server webserver.WebServerInterface,
	uiApp *tview.Application,
	menuPages *tview.Pages,
) *Menu {
	return &Menu{
		reader:    bufio.NewReader(os.Stdin),
		server:    server,
		options:   []string{"Change title", "Update body", "Exit"},
		uiApp:     uiApp,
		menuPages: menuPages,
	}
}

func (m *Menu) CreateMenu() *tview.List {
	list := tview.NewList()
	list.AddItem("Change title", "Press to change the title", '1', func() {
		m.handleChangeTitle()
	}).
		AddItem("Update body", "Press to update the body", '2', func() {
			m.handleChangeBody()
		}).
		AddItem("Exit", "Press to exit", '3', func() {
			m.handleExit()
		})

	return list
}

func (m *Menu) handleChangeTitle() {
	focused := m.uiApp.GetFocus()
	form := tview.NewForm()
	form.AddInputField("New title", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		newTitle := form.GetFormItemByLabel("New title").(*tview.InputField).GetText()
		m.server.UpdateTitle(strings.TrimSpace(newTitle))
		m.menuPages.RemovePage("ChangeTitle")
		m.uiApp.SetFocus(focused)
	})
	form.
		SetBorder(true).
		SetTitle("Enter new title").
		SetTitleAlign(tview.AlignLeft)
	m.menuPages.AddPage("ChangeTitle", form, true, true)
	m.uiApp.SetFocus(form)
}

func (m *Menu) handleChangeBody() {
	focused := m.uiApp.GetFocus()
	form := tview.NewForm()
	form.AddInputField("New body", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		newBody := form.GetFormItemByLabel("New body").(*tview.InputField).GetText()
		m.server.UpdateBody(strings.TrimSpace(newBody))
		m.uiApp.SetFocus(focused)
	})
	form.SetBorder(true).SetTitle("Enter new body").SetTitleAlign(tview.AlignLeft)
	m.menuPages.AddPage("ChangeBody", form, true, true)
	m.uiApp.SetFocus(form)
}

func (m *Menu) handleExit() {
	err := m.server.Stop()
	if err != nil {
		log.Fatalf("Error stopping server: %v", err)
	}
	m.uiApp.Stop()
}

func (m *Menu) GetOptions() []string {
	return m.options
}

func (m *Menu) GetApp() *tview.Application {
	return m.uiApp
}

func (m *Menu) GetPages() *tview.Pages {
	return m.menuPages
}
