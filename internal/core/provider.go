package core

type CommandProvider interface {
	List() []Command
}
