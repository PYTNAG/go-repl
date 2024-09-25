package commands

type Function interface {
	WithArgs(...string) Function
	Execute() Result
}
