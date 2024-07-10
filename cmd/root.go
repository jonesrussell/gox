package cmd

import (
	"bytes"
	"jonesrussell/gocreate/ui"
	"jonesrussell/gocreate/webserver"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	Debug bool
}

func NewConfig() *Config {
	return &Config{}
}

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

var rootCmd = &cobra.Command{
	Use:   "gocreate",
	Short: "Create webs things!",
}

func NewRootCmd(
	server webserver.WebServerInterface,
	menu ui.MenuInterface,
	ui ui.UIInterface,
) *cobra.Command {
	cfg := NewConfig()

	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", false, "Enable debug mode")

	websiteCommand := NewWebsiteCommand(cfg, server, menu, ui)
	describeCommand := NewDescribeCommand(cfg)
	detectCommand := NewDetectCommand(cfg)

	rootCmd.AddCommand(websiteCommand.Command())
	rootCmd.AddCommand(describeCommand.Command())
	rootCmd.AddCommand(detectCommand.Command())

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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
