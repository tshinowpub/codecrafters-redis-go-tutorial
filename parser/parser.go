package parser

import (
	"errors"
	"fmt"
)

const (
	Separator = "\\r\\n"
)

type Parser struct {
}

func New() Parser {
	return Parser{}
}

func (parser *Parser) parseArray(values []byte) ParseResult {
	fmt.Printf("入力値：, %s\n", values)

	respInput, err := newRespInput(values)
	if err != nil {
		return parseFailed(err)
	}

	respInputType := respInput.respType()
	if respInputType.IsUnsupported() {
		return parseFailed(errors.New("unsupported protocol"))
	}

	return respInput.process()
}
