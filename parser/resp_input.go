package parser

import (
	"bytes"
	"errors"
	"strconv"
)

type RespInput struct {
	input []byte
}

func newRespInput(byteString []byte) (*RespInput, error) {
	if len(byteString) == 0 {
		return nil, errors.New("must not be byte string is empty")
	}

	return &RespInput{input: byteString}, nil
}

func (r *RespInput) respType() RespInputType {
	return FromFirstByte(r.input[0])
}

func (r *RespInput) calcArrayLength() (int, error) {
	index := bytes.Index(r.input, []byte(Separator))
	if index == -1 {
		return -1, errors.New("byte string is not resp format")
	}

	length, err := strconv.Atoi(string(r.input[1:index]))
	if err != nil {
		return -1, errors.New("byte string is not resp format")
	}

	return length, nil
}

func (r *RespInput) process() ParseResult {
	count, calcArrayLengthError := r.calcArrayLength()
	if calcArrayLengthError != nil {
		return parseFailed(calcArrayLengthError)
	}

	/**
	 * Resp配列のバイト数を取得
	 */
	firstValueIndex := bytes.Index(r.input, []byte("$"))
	if firstValueIndex == -1 {
		return parseFailed(errors.New("byte string is not resp format in order string"))
	}

	value := r.input[firstValueIndex:]

	var tokens []Token
	for i := 1; i <= count; i++ {
		index := bytes.Index(value, []byte(Separator))
		if index == -1 {
			return parseFailed(errors.New("byte string is not resp format"))
		}

		byteLength, err := strconv.Atoi(string(value[1:index]))
		if err != nil {
			return parseFailed(errors.New("byte string is not resp format"))
		}

		/**
		 * Resp配列の値を取得
		 */
		value = value[index+len([]byte(Separator)):]

		index = bytes.Index(value, []byte(Separator))
		if index == -1 {
			return parseFailed(errors.New("byte string is not resp format, resp array value is not found"))
		}

		respValue := value[0:index]
		if byteLength != len(respValue) {
			return parseFailed(errors.New("byte string is not resp format, The number of bytes specified differs from the actual value"))
		}

		tokens = append(tokens, respValue)

		value = value[index+byteLength:]
	}

	return parseSucceeded(tokens)
}
