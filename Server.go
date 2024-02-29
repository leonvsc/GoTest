package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 8080")

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a new scanner to read from the connection
	scanner := bufio.NewScanner(conn)

	for {
		// Read a line from the connection
		scanner.Scan()
		received := scanner.Text()

		// Print the received message
		fmt.Println("Received:", received)

		// Check if the message is "quit"
		if strings.TrimSpace(received) == "quit" {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}

		// Send a response
		fmt.Print("Reply: ")
		replyScanner := bufio.NewScanner(os.Stdin)
		replyScanner.Scan()
		reply := replyScanner.Text()
		conn.Write([]byte(reply + "\n"))
	}
}
