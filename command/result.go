package command

//go:generate enumer -type=ResultType -output=./result_type_enumer.go
type ResultType int

type Result struct {
	_type   ResultType
	message string
}

const (
	Ok ResultType = iota + 1
	Failed
)

func ok() Result {
	return Result{_type: Ok, message: ""}
}

func failed(message string) Result {
	return Result{message: message}
}

func (r *Result) isOk() bool {
	return r._type == 1
}
