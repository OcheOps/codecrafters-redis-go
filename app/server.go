package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
	SERVER_ADDRESS = "localhost:6379"
)

func parseRequest(request string) (string, []string) {
	// Split the request by newline characters
	tokens := strings.Split(request, "\r\n")

	// The first token should be the "*<number-of-tokens>" pattern
	numTokens := len(tokens) - 1 // Subtract 1 to account for the last empty token

	// Extract the command and arguments
	command := strings.ToLower(strings.TrimPrefix(tokens[1], "$"))
	args := make([]string, 0, numTokens-1)
	for _, arg := range tokens[2 : numTokens+1] {
		args = append(args, strings.TrimPrefix(arg, "$"))
	}

	return command, args
}

func main() {
	// Create a TCP listener
	listener, err := net.Listen("tcp", SERVER_ADDRESS)
	if err != nil {
		fmt.Println("Failed to create listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", SERVER_ADDRESS)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a buffered reader for the connection
	reader := bufio.NewReader(conn)

	for {
		// Read the incoming request
		request, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read request:", err)
			return
		}

		// Parse the request
		command, args := parseRequest(request)

		var response string
		switch command {
		case "echo":
			response = handleEcho(args)
		default:
			response = "-ERR unknown command '" + command + "'\r\n"
		}

		// Send the response
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Failed to send response:", err)
			return
		}
	}
}