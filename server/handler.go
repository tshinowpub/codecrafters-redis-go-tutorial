package server

import (
	"fmt"
	"github.com/tshinowpub/codecrafters-redis-go/parser"
)

func Handle(strings []byte) Response {
	input, _ := parser.NewRespInput(strings)

	switch {
	case input.IsArray() == true:
		respParser := parser.New()

		result := respParser.ParseArray(input)
		if result.IsError() {
			return badRequest([]byte(result.GetErrorMessage()))
		}

		command, _ := result.GetCommand()

		switch command {
		case "ECHO":
			break
		}

	case input.IsUnsupported() == true:
		fmt.Println("Input Value was unsupported protocol.")
	}

	return ok()
}
