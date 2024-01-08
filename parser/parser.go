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

func (r *RespArrayResult) Length() int {
	return len(r.value)
}

func (r *RespArrayResult) Values() [][]byte {
	return r.value
}

func New() Parser {
	return Parser{}
}

func (parser *Parser) ParseArray(values []byte) (*RespArrayResult, error) {
	fmt.Printf("テストだよ！, %s\n", values)

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
