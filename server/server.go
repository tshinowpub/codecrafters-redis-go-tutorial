package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

type Server struct {
	wg           *sync.WaitGroup
	chTerminated chan bool
	listener     net.Listener
	port         uint16
}

func NewServer(wg *sync.WaitGroup, port uint16) *Server {
	return &Server{wg: wg, chTerminated: make(chan bool), port: port}
}

func (s *Server) Listen() {
	defer s.wg.Done()

	fmt.Println("Logs from your program will appear here!")

	fmt.Printf("Socket opened. Address was 0.0.0.0:%s.\n", "6380")

	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(int(s.port)))
	if err != nil {
		fmt.Println("Failed to bind to port " + strconv.Itoa(int(s.port)))
		os.Exit(1)
	}

	s.listener = listener

	defer listener.Close()

	var connWg sync.WaitGroup

	for {
		select {
		case <-s.chTerminated:
			fmt.Println("Received signal. Server shutting down.")
			return
		default:
			conn, err := listener.Accept()
			fmt.Println("Finished Listener Accepted!!")
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				return
			}

			connWg.Add(1)

			go handleConnection(conn, &connWg)
		}
	}
}

func (s *Server) Terminate() {
	if s.listener == nil {
		return
	}

	close(s.chTerminated)

	s.listener.Close()
}

func handleConnection(conn net.Conn, connWg *sync.WaitGroup) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	response := Handle(buffer[:n])

	_, err = conn.Write(response.GetValue())
	if err != nil {
		fmt.Println("Failed write connection: ", err.Error())
		os.Exit(1)
	}

	connWg.Done()
}
