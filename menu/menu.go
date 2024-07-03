package menu

import (
	"bufio"
	"os"
	"strings"

	"jonesrussell/gocreate/websiteserver"

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

type menuImpl struct {
	reader    *bufio.Reader
	server    *websiteserver.WebsiteServerInterface
	options   []string
	uiApp     *tview.Application
	menuPages *tview.Pages
}

// Ensure menuImpl implements MenuInterface
var _ MenuInterface = &menuImpl{}

func NewMenu(
	server *websiteserver.WebsiteServerInterface,
	uiApp *tview.Application,
	menuPages *tview.Pages,
) *menuImpl {
	return &menuImpl{
		reader:    bufio.NewReader(os.Stdin),
		server:    server,
		options:   []string{"Change title", "Update body", "Exit"},
		uiApp:     uiApp,
		menuPages: menuPages,
	}
}

func (m *menuImpl) CreateMenu() *tview.List {
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

func (m *menuImpl) handleChangeTitle() {
	focused := m.uiApp.GetFocus()
	form := tview.NewForm()
	form.AddInputField("New title", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		newTitle := form.GetFormItemByLabel("New title").(*tview.InputField).GetText()
		(*m.server).UpdateTitle(strings.TrimSpace(newTitle))
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

func (m *menuImpl) handleChangeBody() {
	focused := m.uiApp.GetFocus()
	form := tview.NewForm()
	form.AddInputField("New body", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		newBody := form.GetFormItemByLabel("New body").(*tview.InputField).GetText()
		(*m.server).UpdateBody(strings.TrimSpace(newBody))
		m.uiApp.SetFocus(focused)
	})
	form.SetBorder(true).SetTitle("Enter new body").SetTitleAlign(tview.AlignLeft)
	m.menuPages.AddPage("ChangeBody", form, true, true)
	m.uiApp.SetFocus(form)
}

func (m *menuImpl) handleExit() {
	(*m.server).Stop()
	m.uiApp.Stop()
}

func (m *menuImpl) GetOptions() []string {
	return m.options
}

func (m *menuImpl) GetApp() *tview.Application {
	return m.uiApp
}

func (m *menuImpl) GetPages() *tview.Pages {
	return m.menuPages
}
