package command

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEcho(t *testing.T) {
	echo := Echo("Hello World!!")

	res := echo.run()

	require.True(t, res.isOk())
}
