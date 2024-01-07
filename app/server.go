package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running test.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	fmt.Println("Socket opened. Address was 0.0.0.0:6379.")

	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
