// menu/menu.go

package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jonesrussell/gocreate/websiteserver"
)

type MenuInterface interface {
	Display()
	handleChangeTitle() error
	handleExit()
}

type menuImpl struct {
	reader *bufio.Reader
	server *websiteserver.WebsiteServerInterface
}

// Ensure menuImpl implements MenuInterface
var _ MenuInterface = &menuImpl{}

func NewMenu(server *websiteserver.WebsiteServerInterface) *menuImpl {
	return &menuImpl{
		reader: bufio.NewReader(os.Stdin),
		server: server,
	}
}

func (m *menuImpl) Display() {
	for {
		fmt.Println("\nInteractive Menu:")
		fmt.Println("1. Change title")
		fmt.Println("2. Update body")
		fmt.Println("3. Exit")
		fmt.Print("Enter command number: ")

		command, _ := m.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "1":
			err := m.handleChangeTitle()
			if err != nil {
				fmt.Println("Error updating title:", err)
			}
		case "2":
			err := m.handleChangeBody()
			if err != nil {
				fmt.Println("Error updating body:", err)
			}
		case "3":
			m.handleExit()
			return
		default:
			fmt.Println("Invalid command. Please enter a number from 1 to 3.")
		}
	}
}

func (m *menuImpl) handleChangeTitle() error {
	fmt.Print("Enter new title: ")
	newTitle, err := m.reader.ReadString('\n')
	if err != nil {
		return err
	}
	(*m.server).UpdateTitle(strings.TrimSpace(newTitle)) // Remove newline
	return nil
}

func (m *menuImpl) handleChangeBody() error {
	fmt.Print("Enter new body content: ")
	newBody, err := m.reader.ReadString('\n')
	if err != nil {
		return err
	}
	(*m.server).UpdateBody(strings.TrimSpace(newBody)) // Assuming UpdateBody exists
	return nil
}

func (m *menuImpl) handleExit() {
	fmt.Println("Exiting...")
	(*m.server).Stop()
}
