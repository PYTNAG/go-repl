package commands

type Result struct {
	isTerminate bool
	message     string
}

func NewResult(isTerminate bool, message string) *Result {
	result := new(Result)

	result.isTerminate = isTerminate
	result.message = message

	return result
}

func (r *Result) Message() string {
	return r.message
}

func (r *Result) IsTerminate() bool {
	return r.isTerminate
}
