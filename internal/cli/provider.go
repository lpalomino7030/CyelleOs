package cli

type CommandProvider interface {
	List() []Command
}