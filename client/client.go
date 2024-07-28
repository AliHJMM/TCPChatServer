package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Get command line arguments
	args := os.Args
	if len(args) != 3 {
		// If the number of arguments is incorrect, print usage instructions and exit
		fmt.Println("[USAGE]: ./client $port $host")
		return
	}

	// Extract port and host from command line arguments
	port := args[1]
	host := args[2]

	// Attempt to connect to the server at the specified host and port
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		// If connection fails, print error message and exit
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close() // Ensure the connection is closed when main exits

	// Create a buffered reader for the connection
	reader := bufio.NewReader(conn)
	// Print the initial message from the server
	fmt.Print(readMessage(reader))

	// Start a goroutine to continuously read messages from the server
	go readServerMessages(reader)

	// Create a buffered writer for the connection
	writer := bufio.NewWriter(conn)
	for {
		// Read user input from the console
		fmt.Print("[YOU]: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if message == "\n" {
			// Skip empty messages
			continue
		}
		// Send the user's message to the server
		sendMessage(writer, message)
	}
}

// readMessage reads a single line message from the server
func readMessage(reader *bufio.Reader) string {
	message, _ := reader.ReadString('\n')
	return message
}

// sendMessage sends a message to the server
func sendMessage(writer *bufio.Writer, message string) {
	_, err := writer.WriteString(message)
	if err != nil {
		// If sending the message fails, print error message and exit
		fmt.Println("Error sending message:", err.Error())
		os.Exit(1)
	}
	writer.Flush() // Ensure the message is sent immediately
}

// readServerMessages continuously reads messages from the server and prints them
func readServerMessages(reader *bufio.Reader) {
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// If reading the message fails, print error message and stop reading
			fmt.Println("Error reading server message:", err)
			return
		}
		// Print the received message
		fmt.Print(message)
	}
}
