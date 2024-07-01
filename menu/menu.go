// menu/menu.go

package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jonesrussell/gocreate/httpserver"
)

type Menu struct {
	reader *bufio.Reader
	server *httpserver.Server
}

func NewMenu(server *httpserver.Server) *Menu {
	return &Menu{
		reader: bufio.NewReader(os.Stdin),
		server: server,
	}
}

func (m *Menu) Display() {
	for {
		fmt.Println("\nInteractive Menu:")
		fmt.Println("1. Change title")
		fmt.Println("2. Exit")
		fmt.Print("Enter command number: ")

		command, _ := m.reader.ReadString('\n')
		command = strings.TrimSpace(command) // Remove newline

		switch command {
		case "1":
			err := m.handleChangeTitle()
			if err != nil {
				fmt.Println("Error updating title:", err)
			}
		case "2":
			m.handleExit()
			return
		default:
			fmt.Println("Invalid command. Please enter a number from 1 to 2.")
		}
	}
}

func (m *Menu) handleChangeTitle() error {
	fmt.Print("Enter new title: ")
	newTitle, err := m.reader.ReadString('\n')
	if err != nil {
		return err
	}
	m.server.UpdateTitle(strings.TrimSpace(newTitle)) // Remove newline
	return nil
}

func (m *Menu) handleExit() {
	fmt.Println("Exiting...")
	m.server.Stop()
}
