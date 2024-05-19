// handle.go
package main

import (
    "bufio"
    "fmt"
    "io"
    "net"
    "strconv"
    "strings"
)

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