package parser

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

const (
	Separator = "\\r\\n"
)

type Parser struct {
}

func New() Parser {
	return Parser{}
}

func (parser *Parser) parseArray(input *RespInput) ParseResult {
	fmt.Printf("入力値：, %s\n", input.input)

	count, err := parser.calcLength(input)
	if err != nil {
		return parseFailed(err)
	}

	/**
	 * Resp配列のバイト数を取得
	 */
	firstValueIndex := bytes.Index(input.values(), []byte("$"))
	if firstValueIndex == -1 {
		return parseFailed(errors.New("byte string is not resp format in order string"))
	}

	value := input.values()[firstValueIndex:]

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

func (parser *Parser) calcLength(input *RespInput) (int, error) {
	index := bytes.Index(input.values(), []byte(Separator))
	if index == -1 {
		return -1, errors.New("byte string is not resp format")
	}

	length, err := strconv.Atoi(string(input.values()[1:index]))
	if err != nil {
		return -1, errors.New("byte string is not resp format")
	}

	return length, nil
}
