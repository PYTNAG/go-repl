package commands

type Function interface {
	WithArgs(...string)
	Execute() Result
}
