package commands

import (
	"fmt"

	"cyelle/internal/core"
)

type HelpCommand struct {
	provider core.CommandProvider
}

func (HelpCommand) Name() string {
	return "help"
}

func NewHelpCommand(provider core.CommandProvider) HelpCommand {
	return HelpCommand{
		provider: provider,
	}
}

func (HelpCommand) Description() string {
	return "Show available commands."
}

func (h HelpCommand) Run(args []string) int {

	fmt.Println("Cyelle CLI")
	fmt.Println()

	fmt.Println("Available commands:")

	for _, cmd := range h.provider.List() {

		fmt.Printf("  %-10s %s\n",
			cmd.Name(),
			cmd.Description())

	}

	return 0
}
