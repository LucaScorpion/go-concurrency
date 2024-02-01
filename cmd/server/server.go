package main

import (
	"bufio"
	"fmt"
	"net"
)

var port = 7000

func main() {
	fmt.Println("Server listening on port", port)
	l, listenErr := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if listenErr != nil {
		panic(listenErr)
	}

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

	reader := bufio.NewReader(con)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// Use Print instead of Println since line already has a trailing newline.
		fmt.Print("< ", line)
	}

	fmt.Println("Connection closed")
}
