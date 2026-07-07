package commands

import (
	"fmt"
	"os/exec"
)

type CleanCommand struct{}

func (CleanCommand) Name() string {
	return "clean"
}

func (CleanCommand) Description() string {
	return "Clean the workspace CyelleOs ISO."
}

func (CleanCommand) Run(args []string) int {
	fmt.Println("Clean workspace CyelleOs...")
	cmd := exec.Command("rm", "-rf", "work")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error cleaning workspace:", err)
		return 1
	}
	return 0
}
