package commands

type Result interface {
	IsTerminate() bool
	Message() string
}
