package parser

import "errors"

type RespArrayInput struct {
	input []byte
}

func NewRespArrayInput(byteString []byte) (*RespArrayInput, error) {
	if len(byteString) == 0 {
		return nil, errors.New("must not be byte string is empty")
	}

	return &RespArrayInput{input: byteString}, nil
}
