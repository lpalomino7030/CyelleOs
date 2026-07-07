package commands

import "fmt"

type VersionCommand struct{}

func (VersionCommand) Name() string {
	return "version"
}

func (VersionCommand) Description() string {
	return "Version info Cyelle"
}

func (VersionCommand) Run(args []string) int {
	fmt.Println("Version 0.0.1")
	return 0
}