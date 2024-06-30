package cmd

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
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
		http.HandleFunc("/", handleRequest)

		log.Println("Listening on :3000...")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	content, err := readFile("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	changeTitle(doc, "New Title")

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
