package cmd

import (
	"bytes"
	"jonesrussell/gocreate/menu"
	"jonesrussell/gocreate/websiteserver"
	"os"

	"github.com/spf13/cobra"
)

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
	Short: "Create things!",
}

func NewRootCmd(server websiteserver.WebsiteServerInterface, menu menu.MenuInterface) *cobra.Command {
	rootCmd.AddCommand(NewWebsiteCommand(server, menu)) // Use the factory function for the website command
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocreate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
