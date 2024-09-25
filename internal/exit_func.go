package internal

import "github.com/pytnag/go-repl/commands"

const ExitResultText string = "exit"

type ExitFunc struct{}

func (f *ExitFunc) WithArgs(...string) commands.Function {
	return f
}

func (f *ExitFunc) Execute() *commands.Result {
	return commands.NewResult(true, ExitResultText)
}
