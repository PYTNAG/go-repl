package internal

import "github.com/pytnag/go-repl/commands"

type SampleCommand struct{}

func (cmd *SampleCommand) Name() string {
	return "sample"
}

func (cmd *SampleCommand) Aliases() []string {
	return nil
}

func (cmd *SampleCommand) EvaluateFunction() commands.Function {
	return nil
}

func (cmd *SampleCommand) Description() *commands.Description {
	return commands.NewDescription(cmd.Name(), cmd.Aliases()...)
}
