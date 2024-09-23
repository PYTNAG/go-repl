package repl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pytnag/go-repl/commands"
	"github.com/pytnag/go-repl/repl"
)

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

func TestCreatingEmptyRepl(t *testing.T) {
	repl := MakeRepl(t)

	commands := repl.Descriptions()
	assert.NotNil(t, commands)
	assert.Empty(t, commands)
}

func TestCreatingNotEmptyRepl(t *testing.T) {
	repl := MakeRepl(t, &SampleCommand{})

	commands := repl.Descriptions()
	assert.NotNil(t, commands)
	assert.NotEmpty(t, commands)
}

func MakeRepl(t *testing.T, commands ...commands.Command) *repl.Repl {
	replConfig := repl.NewRepl()

	for _, cmd := range commands {
		replConfig.AddCommand(cmd)
	}

	repl := replConfig.Build()

	assert.NotNil(t, repl)
	return repl
}
