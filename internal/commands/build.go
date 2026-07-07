package commands

import (
	"fmt"
	"os"
	"os/exec"
)

type BuildCommand struct{}

func (BuildCommand) Name() string {
	return "build"
}

func (BuildCommand) Description() string {
	return "Build the CyelleOs ISO."
}

func (BuildCommand) Run(args []string) int {
	if _, err := exec.LookPath("mkarchiso"); err != nil {
		fmt.Println("Error: mkarchiso is not installed")
		return 0

	}

	fmt.Println("Building CyelleOs...")
	cmd := exec.Command(
		"mkarchiso",
		"-v",
		"-w", "build/work",
		"-o", "build/out",
		"profile",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("📦 Building CyelleOS ISO...")

	if err := cmd.Run(); err != nil {
		return 0
	}

	fmt.Println("✅ ISO generated successfully!")

	return 0
}
