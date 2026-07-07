package commands

import "fmt"

type DoctorCommand struct{}

func (DoctorCommand) Name() string {
	return "doctor"
}

func (DoctorCommand) Description() string {
	return "Check the development environment."
}

func (DoctorCommand) Run(args []string) int {
	fmt.Println("Checking environment...")
	return 0
}