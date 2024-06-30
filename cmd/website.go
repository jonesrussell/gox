package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

var (
	serverMux *http.ServeMux
	server    *http.Server
	wg        sync.WaitGroup
)

// websiteCmd represents the website command
var websiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Create a no code website from the command line",
	Long: `The 'website' command allows you to create a
	no-code website directly from the command line. This is
	particularly useful for beginner developers who need to
	quickly	set up a static website.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("\nInteractive Menu:")
			fmt.Println("1. Change title")
			fmt.Println("2. Exit")
			fmt.Print("Enter command number: ")

			command, _ := reader.ReadString('\n')
			command = strings.TrimSpace(command) // Remove newline

			switch command {
			case "1":
				newTitle := promptForTitle()

				serverMux = http.NewServeMux()
				serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					handleRequest(w, r, newTitle)
				})

				if server != nil {
					if err := server.Close(); err != nil {
						log.Fatal(err)
					}
					wg.Wait() // Wait for the server to shutdown
				}

				server = &http.Server{Addr: ":3000", Handler: serverMux}

				wg.Add(1)
				go func() {
					defer wg.Done()

					log.Println("Listening on :3000...")
					if err := server.ListenAndServe(); err != http.ErrServerClosed {
						log.Fatal(err)
					}
				}()
			case "2":
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid command. Please enter a number from 1 to 2.")
			}
		}
	},
}

func promptForTitle() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter new title: ")
	newTitle, _ := reader.ReadString('\n')
	return strings.TrimSpace(newTitle) // Remove newline
}

func handleRequest(w http.ResponseWriter, r *http.Request, newTitle string) {
	content, err := readFile("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	changeTitle(doc, newTitle)

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		log.Fatal(err)
	}

	http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(buf.Bytes()))
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func changeTitle(n *html.Node, newTitle string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			n.FirstChild.Data = newTitle
		}
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		changeTitle(c, newTitle)
	}
}

func init() {
	rootCmd.AddCommand(websiteCmd)
}
