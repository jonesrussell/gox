package cmd

import (
	"bytes"
	"fmt"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"os"

	"github.com/spf13/cobra"
)

var Debug bool // Global variable for debug mode

// ExecuteCommandC is a helper function that will execute the Cobra command
// and return the command, output, and any errors that occurred.
func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	fmt.Println("In ExecuteCommandC")
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	fmt.Println("before root.ExecuteC()")
	c, err = root.ExecuteC()
	fmt.Println("after root.ExecuteC()")
	return c, buf.String(), err
}

// ExecuteCommand is a wrapper around ExecuteCommandC that only returns the output and error.
func ExecuteCommand(root *cobra.Command, args ...string) (output string, err error) {
	fmt.Println("In ExecuteCommand")
	_, output, err = ExecuteCommandC(root, args...)
	fmt.Println("after ExecuteCommandC")
	fmt.Println(output)
	return output, err
}

// In your root.go file, update the NewRootCmd function to include the website command

var rootCmd = &cobra.Command{
	Use:   "gocreate",
	Short: "Create webs things!",
}

func NewRootCmd(
	server webserver.WebServerInterface,
	menu ui.MenuInterface,
	ui ui.UIInterface,
) *cobra.Command {
	websiteCommand := NewWebsiteCommand(server, menu, ui)
	rootCmd.AddCommand(websiteCommand.Command())
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Enable debug mode")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocreate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
