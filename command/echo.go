package command

import (
	"fmt"
)

type Echo []byte

func (e Echo) run() Result {
	fmt.Println("Sending", string(e))

	return ok()
}
