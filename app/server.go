package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//Uncomment this block to pass the first stage
	
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

	go func(c net.Conn) {
		defer c.Close()
		for {
			buffer := make([]byte, 1024)
			length, err := c.Read(buffer)
			if err != nil {
				if err != io.EOF {
					fmt.Println("Read error:", err)
				}
				break
			}
	
			command := strings.TrimSpace(string(buffer[:length]))
			if command == "PING" {
				c.Write([]byte("+PONG\r\n"))
			} else if strings.HasPrefix(command, "ECHO") {
				echoText := strings.TrimPrefix(command, "ECHO ")
				c.Write([]byte("$" + strconv.Itoa(len(echoText)) + "\r\n" + echoText + "\r\n"))
			}
		}
	}(conn)
}
}


	


// func respondToPing(conn net.Conn) {
// 	_, err := conn.Write([]byte("+PONG\r\n"))
// 	if err != nil {
// 		log.Println("Failed to respond to PING:", err)
// 	}
// }

