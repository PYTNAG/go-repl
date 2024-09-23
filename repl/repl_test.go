package repl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pytnag/go-repl/repl"
)

func TestCreatingEmptyRepl(t *testing.T) {
	repl := repl.NewRepl().Build()
	assert.NotNil(t, repl)

	commands := repl.Descriptions()
	assert.NotNil(t, commands)
	assert.Empty(t, commands)
}
