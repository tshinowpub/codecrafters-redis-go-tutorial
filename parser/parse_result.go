package parser

import "errors"

type Token []byte

type ParseResult struct {
	tokens     []Token
	parseError error
}

func (p *ParseResult) IsError() bool {
	return p.parseError != nil
}

func (p *ParseResult) GetCommand() (string, error) {
	if p.parseError != nil || len(p.tokens) == 0 {
		return "", errors.New("no command")
	}

	return string(p.tokens[0]), nil
}

func (p *ParseResult) GetErrorMessage() string {
	return p.parseError.Error()
}

func parseSucceeded(tokens []Token) ParseResult {
	return ParseResult{tokens, nil}
}

func parseFailed(parseError error) ParseResult {
	return ParseResult{parseError: parseError}
}

func (p *ParseResult) length() int {
	return len(p.tokens)
}

func (p *ParseResult) values() []Token {
	return p.tokens
}
