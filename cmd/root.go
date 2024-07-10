package cmd

import (
	"bytes"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"os"

	"github.com/spf13/cobra"
)

var Debug bool // Global variable for debug mode

// ExecuteCommandC is a helper function that will execute the Cobra command
// and return the command, output, and any errors that occurred.
func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	return c, buf.String(), err
}

// ExecuteCommand is a wrapper around ExecuteCommandC that only returns the output and error.
func ExecuteCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = ExecuteCommandC(root, args...)
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
	symexCommand := NewSymexCommand() // Assuming you have a similar constructor for SymexCommand

	rootCmd.AddCommand(websiteCommand.Command())
	rootCmd.AddCommand(symexCommand.Command())

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
