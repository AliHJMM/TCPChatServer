package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("[USAGE]: ./client $port $host")
		return
	}

	port := args[1]
	host := args[2]

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Print(readMessage(reader))

	go readServerMessages(reader)

	writer := bufio.NewWriter(conn)
	for {
		fmt.Print("[YOU]: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if message == "\n" {
			continue
		}
		sendMessage(writer, message)
	}
}


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
