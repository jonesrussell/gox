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
	Display() *tview.List
	handleChangeTitle()
	handleChangeBody()
	handleExit()
	GetOptions() []string
	GetApp() *tview.Application
	GetPages() *tview.Pages
}

type menuImpl struct {
	reader  *bufio.Reader
	server  *websiteserver.WebsiteServerInterface
	options []string
	uiApp   *tview.Application
	uiPages *tview.Pages
}

// Ensure menuImpl implements MenuInterface
var _ MenuInterface = &menuImpl{}

func NewMenu(
	server *websiteserver.WebsiteServerInterface,
	uiApp *tview.Application,
	uiPages *tview.Pages,
) *menuImpl {
	return &menuImpl{
		reader:  bufio.NewReader(os.Stdin),
		server:  server,
		options: []string{"Change title", "Update body", "Exit"},
		uiApp:   uiApp,
		uiPages: uiPages,
	}
}

func (m *menuImpl) Display() *tview.List {
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
	form := tview.NewForm()
	form.AddInputField("New title", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		newTitle := form.GetFormItemByLabel("New title").(*tview.InputField).GetText()
		(*m.server).UpdateTitle(strings.TrimSpace(newTitle))
		m.uiApp.SetFocus(m.uiPages)
	})
	form.SetBorder(true).SetTitle("Enter new title").SetTitleAlign(tview.AlignLeft)
	m.uiPages.AddPage("ChangeTitle", form, true, true)
	m.uiApp.SetFocus(form)
}

func (m *menuImpl) handleChangeBody() {
	form := tview.NewForm()
	form.AddInputField("New body", "", 20, nil, nil)
	form.AddButton("Submit", func() {
		newBody := form.GetFormItemByLabel("New body").(*tview.InputField).GetText()
		(*m.server).UpdateBody(strings.TrimSpace(newBody))
		m.uiApp.SetFocus(m.uiPages)
	})
	form.SetBorder(true).SetTitle("Enter new body").SetTitleAlign(tview.AlignLeft)
	m.uiPages.AddPage("ChangeBody", form, true, true)
	m.uiApp.SetFocus(form)
}

func (m *menuImpl) handleExit() {
	fmt.Println("Exiting...")
	(*m.server).Stop()
}

func (m *menuImpl) GetOptions() []string {
	return m.options
}

func (m *menuImpl) GetApp() *tview.Application {
	return m.uiApp
}

func (m *menuImpl) GetPages() *tview.Pages {
	return m.uiPages
}
