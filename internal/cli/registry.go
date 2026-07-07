package cli

import "cyelle/internal/core"

type Registry struct {
	commands map[string]core.Command
}

func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]core.Command),
	}
}

func (r *Registry) List() []core.Command {

	commands := make([]core.Command, 0, len(r.commands))

	for _, command := range r.commands {
		commands = append(commands, command)
	}

	return commands
}

func (r *Registry) Register(command core.Command) {
	r.commands[command.Name()] = command
}

func (r *Registry) Get(name string) (core.Command, bool) {
	cmd, ok := r.commands[name]
	return cmd, ok
}
