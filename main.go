package main

import (
	"TcpChat/server"
	"fmt"
	"log"
	"os"
)

func main(){
	var port string
	args := os.Args
	// Set default port if no argument is provided
	if len(args) == 1 {
		port = ":8989"
	} else if len(args) == 2 {
		port = ":" + args[1]
	} else {
		// Print usage instructions and exit if more than one argument is provided
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	// Initialize the server with the specified port and a maximum of 10 connections
	server := server.InitServer(port, 10)
	// Start broadcasting messages in a separate goroutine
	go server.BroadcastMessages()
	// Print the server's IP address and port
	fmt.Println("Sever Running on: ", server.GetIp(), "And Port", port)
	// Run the server and log any fatal errors
	log.Fatal(server.RunServer())
}