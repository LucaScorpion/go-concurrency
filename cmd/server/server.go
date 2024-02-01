package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var port = 7000

var connections []net.Conn

func main() {
	// Start the TCP server.
	fmt.Println("Server listening on port", port)
	l, listenErr := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if listenErr != nil {
		panic(listenErr)
	}

	acceptConnections(l)

	// Read from stdin, send lines to all connections.
	stdin := bufio.NewReader(os.Stdin)
	for {
		if line, err := stdin.ReadString('\n'); err == nil {
			broadcast(line)
		}
	}
}

func broadcast(msg string) {
	for _, con := range connections {
		con.Write([]byte(msg))
	}
}

func acceptConnections(l net.Listener) {
	// Accept new connections, forever.
	for {
		if newCon, err := l.Accept(); err != nil {
			fmt.Println("Error while accepting connection:", err)
		} else {
			handleConnection(newCon)
		}
	}
}

func handleConnection(con net.Conn) {
	fmt.Println("New connection")
	connections = append(connections, con)
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

		// Use Print instead of Println since line already has a trailing newline.
		fmt.Print("< ", line)
		broadcast(line)
	}
}
