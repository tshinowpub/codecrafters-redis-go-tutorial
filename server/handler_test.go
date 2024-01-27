package server

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandleOk(t *testing.T) {
	response := Handle([]byte("*2\\r\\n$4\\r\\nECHO\\r\\n$3\\r\\nhey\\r\\n"))

	require.True(t, response.isOk())
}

func TestHandleArray(t *testing.T) {
	response := Handle([]byte("*2"))

	require.True(t, response.isError())
}
