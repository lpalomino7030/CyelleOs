package commands

import "fmt"

type BuildCommand struct{}

func (BuildCommand) Name() string {
	return "build"
}

func (BuildCommand) Description() string {
	return "Build the CyelleOs ISO."
}

func (BuildCommand) Run(args []string) int {
	fmt.Println("Building CyelleOs...")
	return 0
}