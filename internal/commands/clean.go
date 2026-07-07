package commands

import "fmt"

type CleanCommand struct{}

func (CleanCommand) Name() string {
	return "clean"
}

func (CleanCommand) Description() string {
	return "Clean the workspace CyelleOs ISO."
}

func (CleanCommand) Run(args []string) int {
	fmt.Println("Clean workspace CyelleOs...")
	return 0
}