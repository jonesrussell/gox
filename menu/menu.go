// menu/menu.go

package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Menu struct {
	reader *bufio.Reader
}

func NewMenu() *Menu {
	return &Menu{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (m *Menu) Display() string {
	fmt.Println("\nInteractive Menu:")
	fmt.Println("1. Change title")
	fmt.Println("2. Add tags and components")
	fmt.Println("3. Reload server")
	fmt.Println("4. Exit")
	fmt.Print("Enter command number: ")

	command, _ := m.reader.ReadString('\n')
	return strings.TrimSpace(command) // Remove newline
}
