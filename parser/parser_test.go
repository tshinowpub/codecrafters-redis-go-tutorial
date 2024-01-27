package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseArrayString(t *testing.T) {
	parser := New()

	input, _ := NewRespInput([]byte("*2\\r\\n$4\\r\\nECHO\\r\\n$3\\r\\nhey\\r\\n"))

	result := parser.ParseArray(input)

	require.Equal(t, 2, result.length())
	require.Equal(t, "ECHO", string(result.values()[0]))
	require.Equal(t, "hey", string(result.values()[1]))
}

func TestParseArrayStringError(t *testing.T) {
	parser := New()

	input, _ := NewRespInput([]byte("*2"))

	result := parser.ParseArray(input)

	require.Equal(t, 0, result.length())
}
