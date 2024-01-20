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

type RespArrayResult struct {
	value [][]byte
}

func NewRespArrayResult() *RespArrayResult {
	return &RespArrayResult{}
}

func (r *RespArrayResult) addValue(value []byte) {
	r.value = append(r.value, value)
}

func (r *RespArrayResult) length() int {
	return len(r.value)
}

func (r *RespArrayResult) values() [][]byte {
	return r.value
}

func (r *RespArrayResult) Command() []byte {
	return r.value[0]
}

func New() Parser {
	return Parser{}
}

func (parser *Parser) parseArray(values []byte) (*RespArrayResult, error) {
	fmt.Printf("入力値：, %s\n", values)

	respInput, err := newRespInput(values)
	if err != nil {
		return NewRespArrayResult(), err
	}

	respInputType := respInput.respType()
	if respInputType.IsUnsupported() {
		return NewRespArrayResult(), errors.New("unsupported protocol")
	}

	result, err := respInput.process()

	return result, nil
}
