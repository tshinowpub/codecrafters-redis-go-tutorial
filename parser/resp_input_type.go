// Code generated .* DO NOT EDIT\.$

package parser

//go:generate enumer -type=RespInputType -json -output=./resp_input_type_enumer.go
type RespInputType int

const (
	SimpleString RespInputType = iota + 1
	Error
	Integer
	BulkString
	Array
	UNSUPPORTED
)

func (i RespInputType) isUnsupported() bool {
	return i.String() == "UNSUPPORTED"
}

func (i RespInputType) isArray() bool {
	return i.String() == "Array"
}

func fromFirstByte(value byte) RespInputType {
	switch value {
	case '*':
		return Array
	default:
		return UNSUPPORTED
	}
}
