package main

import (
	"bufio"
	"fmt"
	"net"
)

var port = 7000

var connections []net.Conn

var incomingMessages = make(chan string)

func main() {
	// Start the TCP server.
	fmt.Println("Server listening on port", port)
	l, listenErr := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if listenErr != nil {
		panic(listenErr)
	}

	// This needs to run in a goroutine, since it loops forever.
	go acceptConnections(l)

	// Broadcast incoming messages to all connections.
	for msg := range incomingMessages {
		broadcast(msg)
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
			// This needs to run in a goroutine, since it will keep reading input from the connection.
			go handleConnection(newCon)
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

		fmt.Print("< ", line)
		incomingMessages <- line
	}
}
