package main

import "fmt"

func handleEcho(args []string) string {
	if len(args) != 1 {
		return "-ERR wrong number of arguments for 'echo' command\r\n"
	}

	return fmt.Sprintf("$%d\r\n%s\r\n", len(args[0]), args[0])
}