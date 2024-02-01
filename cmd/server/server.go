package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var port = 7000

var connections []net.Conn
var connLock = sync.RWMutex{}

var incomingMessages = make(chan Message)

type Message struct {
	msg string
	src net.Conn
}

func main() {
	// Start the TCP server.
	fmt.Println("Server listening on port", port)
	l, listenErr := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if listenErr != nil {
		panic(listenErr)
	}

	// This needs to run in a goroutine, since it loops forever.
	go acceptConnections(l)
	go readFromStdin()

	// Broadcast incoming messages to all connections.
	for msg := range incomingMessages {
		broadcast(msg)
	}
}

func readFromStdin() {
	// Read lines from stdin, treat those as an incoming message.
	stdin := bufio.NewReader(os.Stdin)
	for {
		if line, err := stdin.ReadString('\n'); err == nil {
			incomingMessages <- Message{
				msg: line,
				src: nil,
			}
		}
	}
}

func broadcast(msg Message) {
	// Here we only need read access to the connections, so ensure nothing is writing to it.
	connLock.RLock()
	for _, con := range connections {
		if con != msg.src {
			con.Write([]byte(msg.msg))
		}
	}
	connLock.RUnlock()
}

func acceptConnections(l net.Listener) {
	// Accept new connections, forever.
	for {
		if newCon, err := l.Accept(); err != nil {
			fmt.Println("Error while accepting connection:", err)
		} else {
			// This needs to run in a goroutine, since it will keep reading input from the connection.
			go handleConnection(newCon)
		}
	}
}

func handleConnection(con net.Conn) {
	fmt.Println("New connection")

	// Here we are writing to the connections, so acquire the write lock.
	connLock.Lock()
	connections = append(connections, con)
	connLock.Unlock()

	readFromConnection(con)
	// When we get here, it means reading has stopped, and thus the connection was closed.
	fmt.Println("Connection closed")
}

func readFromConnection(con net.Conn) {
	reader := bufio.NewReader(con)

	for {
		// Keep reading lines. If an error occurs, that means the connection was closed, so we can stop.
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Print("< ", line)
		incomingMessages <- Message{
			msg: line,
			src: con,
		}
	}
}
