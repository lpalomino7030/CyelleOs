package cli

type Registry struct {
	commands map[string]Command
}

func NewRegistry() *Registry{
	return &Registry{
		commands: make(map[string]Command),
	}
}

func (r *Registry) List() []Command {

	commands := make([]Command, 0, len(r.commands))

	for _, command := range r.commands {
		commands = append(commands, command)
	}

	return commands
}

func (r *Registry) Register(command Command) {
	r.commands[command.Name()] = command
}

func (r *Registry) Get(name string) (Command, bool) {
	cmd, ok := r.commands[name]
	return cmd, ok
}

