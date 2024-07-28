package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readServerMessages(reader *bufio.Reader) {
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server message:", err)
			return
		}
		fmt.Print(message)
	}
}
