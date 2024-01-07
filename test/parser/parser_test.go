package parser

import (
	"github.com/stretchr/testify/require"
	"github.com/tshinowpub/codecrafters-redis-go/parser"
	"testing"
)

func TestParseArrayString(t *testing.T) {
	redisParser := parser.New()

	result, err := redisParser.ParseArray([]byte("*2\\r\\n$4\\r\\nECHO\\r\\n$3\\r\\nhey\\r\\n"))

	require.NoError(t, err)
	require.Equal(t, 2, result.Length())
	require.Equal(t, "ECHO", string(result.Values()[0]))
	require.Equal(t, "hey", string(result.Values()[1]))
}
