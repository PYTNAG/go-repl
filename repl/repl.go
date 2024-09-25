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
		if cmd, ok := repl.commands[tokens[0]]; ok {
			result := cmd.EvaluateFunction().WithArgs(tokens[1:]...).Execute()

			print(writer, result.Message())

			if result.IsTerminate() {
				return
			}
		} else {
			print(writer, "Unknown command")
		}
	}
}

func read(reader io.Reader) (input string) {
	fmt.Fscanln(reader, &input)
	return
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
