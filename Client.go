package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// Create a new scanner to read from the standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Create a new scanner to read from the connection
	connScanner := bufio.NewScanner(conn)

	// Start reading from the connection in a separate goroutine
	go func() {
		for connScanner.Scan() {
			fmt.Println("Server:", connScanner.Text())
		}
	}()

	for {
		// Read a line from the standard input
		scanner.Scan()
		text := scanner.Text()

		// Send the text to the server
		conn.Write([]byte(text + "\n"))

		// Check if the message is "quit"
		if strings.TrimSpace(text) == "quit" {
			return
		}
	}
}
