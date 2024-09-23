package repl

import (
	"github.com/pytnag/go-repl/commands"
)

type ReplConfiurator struct {
	commands []commands.Command
}

func NewRepl() *ReplConfiurator {
	config := new(ReplConfiurator)
	return config
}

func (config *ReplConfiurator) AddCommand(cmd commands.Command) {
	config.commands = append(config.commands, cmd)
}

func (config *ReplConfiurator) Build() *Repl {
	repl := new(Repl)

	for _, cmd := range config.commands {
		repl.commands[cmd.Name()] = cmd.EvaluateFunction()

		for _, alias := range cmd.Aliases() {
			repl.commands[alias] = cmd.EvaluateFunction()
		}
	}

	return repl
}
