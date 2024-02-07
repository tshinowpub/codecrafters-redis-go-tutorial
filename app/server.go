package main

import (
	"fmt"
	"github.com/tshinowpub/codecrafters-redis-go/server"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	s := server.NewServer(&wg, 6380)

	go s.Listen()

	<-signalChannel

	s.Terminate()

	wg.Wait()

	fmt.Println("Server shutting down.")
}
