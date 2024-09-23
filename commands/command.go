package commands

type Command interface {
	Name() string
	Aliases() []string
	EvaluateFunction() Function
	Description() *Description
}
