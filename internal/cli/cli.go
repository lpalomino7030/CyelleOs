package cli

import (
	"cyelle/internal/commands"
	"fmt"
)

type CLI struct {
	registry *Registry
}

func New() *CLI {

	registry := NewRegistry()
	help := commands.NewHelpCommand(registry)

	registry.Register(help)

	registry.Register(commands.BuildCommand{})

	registry.Register(commands.CleanCommand{})

	registry.Register(commands.DoctorCommand{})

	registry.Register(commands.VersionCommand{})

	return &CLI{
		registry: registry,
	}

}

func (c *CLI) Run(args []string) int {

	if len(args) < 2 {

		cmd, _ := c.registry.Get("help")

		return cmd.Run(nil)

	}

	cmd, ok := c.registry.Get(args[1])

	if !ok {

		fmt.Println("Unknown command:", args[1])
		fmt.Println()

		help, _ := c.registry.Get("help")

		return help.Run(nil)

	}

	return cmd.Run(args[2:])

}
