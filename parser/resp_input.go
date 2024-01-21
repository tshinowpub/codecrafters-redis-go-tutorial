package parser

import (
	"errors"
)

type RespInput struct {
	input []byte
}

func newRespInput(byteString []byte) (*RespInput, error) {
	if len(byteString) == 0 {
		return nil, errors.New("must not be byte string is empty")
	}

	respType := FromFirstByte(byteString[0])
	if respType.isUnsupported() {
		return nil, errors.New("must not be byte string is empty")
	}

	return &RespInput{input: byteString}, nil
}

func (r *RespInput) values() []byte {
	return r.input
}
