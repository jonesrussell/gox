package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type SymexCommand struct {
	// Add fields here similar to WebsiteCommand
}

func NewSymexCommand() *SymexCommand {
	// Initialize your SymexCommand here
	cmd := &SymexCommand{
		// Initialize fields here
	}

	return cmd
}

func (s *SymexCommand) HandleDebugFlag(flagset *pflag.FlagSet) bool {
	debug, err := flagset.GetBool("debug")
	if err != nil {
		log.Println("Can't get debug flag, defaulting to false")
		debug = false
	}

	return debug
}

// Add other methods similar to WebsiteCommand

func (s *SymexCommand) Command() *cobra.Command {
	symexCmd := &cobra.Command{
		Use:   "symex",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			debug := s.HandleDebugFlag(cmd.Flags())

			if debug {
				log.Println("Debugging")
			}

			// Add your command's functionality here

			fmt.Println("symex called")
		},
	}

	return symexCmd
}
