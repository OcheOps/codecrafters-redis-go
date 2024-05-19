package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)


func main() {
    l, err := net.Listen("tcp", "0.0.0.0:6379")
    if err != nil {
        fmt.Println("Failed to bind to port 6379")
        os.Exit(1)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection: ", err.Error())
            os.Exit(1)
        }

        go handleConnection(conn)
    }
}
// Remove the duplicate function declaration of handleConnection
// Keep only one instance of handleConnection function
// The code block should contain only the following code:
func handleConnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read error:", err)
			}
			break
		}

		parts := strings.Split(strings.TrimSpace(line), "\r\n")
		if len(parts) < 4 {
			continue
		}

		command := parts[1]
		if command == "$4\r\nPING" {
			c.Write([]byte("+PONG\r\n"))
		} else if command == "$4\r\nECHO" {
			echoText := parts[3]
			c.Write([]byte("$" + strconv.Itoa(len(echoText)) + "\r\n" + echoText + "\r\n"))
		}
	}
}