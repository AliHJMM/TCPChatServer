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
	if len(args) == 1 {
		port = ":8989"
	} else if len(args) == 2 {
		port = ":" + args[1]
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	server := server.InitServer(port, 10) // Set max connections to 10
	go server.BroadcastMessages()
	fmt.Println("Sever Running on: ", server.GetIp(), "And Port", port)
	log.Fatal(server.RunServer())
}