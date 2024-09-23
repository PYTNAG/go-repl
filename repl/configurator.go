package repl

import (
	"github.com/pytnag/go-repl/commands"
)

type ReplConfiurator struct {
	commands []commands.Command
}

func NewRepl() *ReplConfiurator {
	config := new(ReplConfiurator)

	config.commands = make([]commands.Command, 0)

	return config
}

func (config *ReplConfiurator) AddCommand(cmd commands.Command) {
	config.commands = append(config.commands, cmd)
}

func (config *ReplConfiurator) Build() *Repl {
	repl := new(Repl)

	repl.commands = make(map[string]commands.Command)

	for _, cmd := range config.commands {
		repl.commands[cmd.Name()] = cmd

		for _, alias := range cmd.Aliases() {
			repl.commands[alias] = cmd
		}
	}

	return repl
}
