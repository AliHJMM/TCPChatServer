package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readMessage(reader *bufio.Reader) string {
	message, _ := reader.ReadString('\n')
	return message
}

func sendMessage(writer *bufio.Writer, message string) {
	_, err := writer.WriteString(message)
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
		os.Exit(1)
	}
	writer.Flush()
}

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
