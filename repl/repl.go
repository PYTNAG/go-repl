package repl

import (
	"fmt"
	"io"
	"strings"

	"github.com/pytnag/go-repl/commands"
)

type Repl struct {
	commands map[string]commands.Command
}

func (repl *Repl) Begin(reader io.Reader, writer io.Writer) {
	for {
		input := read(reader)

		tokens := strings.Split(input, "")
		result := eval(tokens[0], tokens[1:]...)

		print(writer, result.Message())

		if result.IsTerminate() {
			return
		}
	}
}

func read(reader io.Reader) (input string) {
	fmt.Fscanln(reader, &input)
	return
}

func eval(cmdName string, args ...string) (result commands.Result) {
	return nil
}

func print(writer io.Writer, output string) {
	fmt.Fprint(writer, output)
}

func (repl *Repl) Descriptions() []*commands.Description {
	descriptions := make([]*commands.Description, 0)

	for _, cmd := range repl.commands {
		descriptions = append(descriptions, cmd.Description())
	}

	return descriptions
}
