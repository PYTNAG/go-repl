package commands

type Command interface {
	Name() string
	Aliases() []string
	Evaluate() Result
}
