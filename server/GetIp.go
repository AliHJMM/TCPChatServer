package server

import (
	"log"
	"net"
)

func (s *Server) GetIp() string {
	// Establish a UDP connection to a public server to determine the local IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err) // Log and exit if there's an error
	}
	defer conn.Close()

	// Extract the local address and cast it to a UDPAddr
	localAddress := conn.LocalAddr().(*net.UDPAddr)

	// Return the local IP address as a string
	return localAddress.IP.String()
}
