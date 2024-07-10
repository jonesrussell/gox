package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/mod/modfile"
)

type DetectCommand struct {
	showAllDeps bool
}

func NewDetectCommand() *DetectCommand {
	cmd := &DetectCommand{}
	return cmd
}

func (d *DetectCommand) HandleDebugFlag(flagset *pflag.FlagSet) bool {
	debug, err := flagset.GetBool("debug")
	if err != nil {
		log.Println("Can't get debug flag, defaulting to false")
		debug = false
	}

	return debug
}

func (d *DetectCommand) Command() *cobra.Command {
	detectCmd := &cobra.Command{
		Use:   "detect [path]",
		Short: "Detect if a file or directory is part of a Go project and report basic info",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			debug := d.HandleDebugFlag(cmd.Flags())

			if debug {
				log.Println("Debugging")
			}

			path := args[0]

			if !d.isValidPath(path) {
				fmt.Printf("Invalid path: '%s'\n", path)
				return
			}

			modulePath, dependencies, err := d.detectGoProject(path)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}

			d.printProjectDetails(path, modulePath, dependencies)
		},
	}

	detectCmd.Flags().BoolVarP(&d.showAllDeps, "all-deps", "a", false, "Show all dependencies")

	return detectCmd
}

type dependency struct {
	Name     string
	Version  string
	Indirect bool
}

func (d *DetectCommand) detectGoProject(path string) (string, []dependency, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", nil, fmt.Errorf("error accessing path '%s': %v", path, err)
	}

	if info.Mode().IsRegular() {
		path = filepath.Dir(path)
	}

	for {
		goModPath := filepath.Join(path, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			modulePath, dependencies, err := d.readGoModFile(goModPath)
			if err != nil {
				return "", nil, fmt.Errorf("error reading go.mod file: %v", err)
			}
			return modulePath, dependencies, nil
		}

		parent := filepath.Dir(path)
		if parent == path {
			break
		}
		path = parent
	}

	return "", nil, fmt.Errorf("'%s' is not a Go project. No go.mod file found", path)
}

func (d *DetectCommand) printProjectDetails(projectPath, modulePath string, dependencies []dependency) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.SetStyle(table.StyleColoredBlackOnGreenWhite)

	t.AppendHeader(table.Row{"  Project Name:", filepath.Base(projectPath)})
	t.AppendHeader(table.Row{"  Module Path:", modulePath})
	t.AppendHeader(table.Row{"  Type:", "Go project"})

	t.SetStyle(table.StyleColoredGreenWhiteOnBlack)

	t.AppendSeparator()
	t.AppendHeader(table.Row{"#", "Dependency", "Version", "Indirect"})

	for i, dep := range dependencies {
		t.AppendRow([]interface{}{i + 1, dep.Name, dep.Version, dep.Indirect})
	}

	t.Render()
}

func (d *DetectCommand) readGoModFile(goModPath string) (string, []dependency, error) {
	file, err := os.Open(goModPath)
	if err != nil {
		return "", nil, fmt.Errorf("error opening go.mod file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", nil, fmt.Errorf("error reading go.mod file: %v", err)
	}
	modFile, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing go.mod file: %v", err)
	}

	modulePath := modFile.Module.Mod.Path

	var dependencies []dependency
	dependencies = append(dependencies, d.parseRequire(modFile.Require...)...)

	return modulePath, dependencies, nil
}

func (d *DetectCommand) parseRequire(requires ...*modfile.Require) []dependency {
	var dependencies []dependency
	for _, require := range requires {
		dependencies = append(dependencies, dependency{
			Name:     require.Mod.Path,
			Version:  require.Mod.Version,
			Indirect: require.Indirect,
		})
	}
	return dependencies
}

func (d *DetectCommand) isValidPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
