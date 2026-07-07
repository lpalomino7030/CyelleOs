package commands

import (
	"fmt"

	"cyelle/internal/cli"
)

type HelpCommand struct {
	provider cli.CommandProvider
}

func (HelpCommand) Name() string {
	return "help"
}

func NewHelpCommand(provider cli.CommandProvider) HelpCommand {
	return HelpCommand{
		provider: provider,
	}
}

func (HelpCommand) Description() string {
	return "Show available commands."
}

func (HelpCommand) Run(args []string) int {

	fmt.Println("Cyelle CLI")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  build")
	fmt.Println("  doctor")
	fmt.Println("  clean")
	fmt.Println("  version")
	fmt.Println("  help")

	return 0
}