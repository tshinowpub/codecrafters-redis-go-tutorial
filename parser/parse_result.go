package parser

type Token []byte

type ParseResult struct {
	tokens     []Token
	parseError error
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
