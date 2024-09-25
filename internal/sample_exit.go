package internal

import "github.com/pytnag/go-repl/commands"

const ExitCommandName string = "exit"

type SampleExit struct{}

func (cmd *SampleExit) Name() string {
	return ExitCommandName
}

func (cmd *SampleExit) Aliases() []string {
	return nil
}

func (cmd *SampleExit) EvaluateFunction() commands.Function {
	return &ExitFunc{}
}

func (cmd *SampleExit) Description() *commands.Description {
	return commands.NewDescription(cmd.Name(), cmd.Aliases()...)
}
