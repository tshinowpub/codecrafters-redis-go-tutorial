package parser

import (
	"errors"
)

type RespInput struct {
	input []byte
}

func NewRespInput(byteString []byte) (*RespInput, error) {
	if len(byteString) == 0 {
		return nil, errors.New("must not be byte string is empty")
	}

	respType := fromFirstByte(byteString[0])
	if respType.isUnsupported() {
		return nil, errors.New("must not be byte string is empty")
	}

	return &RespInput{input: byteString}, nil
}

func (r *RespInput) IsArray() bool {
	return fromFirstByte(r.input[0]).isArray()
}

func (r *RespInput) IsUnsupported() bool {
	return fromFirstByte(r.input[0]).isUnsupported()
}

func (r *RespInput) values() []byte {
	return r.input
}
