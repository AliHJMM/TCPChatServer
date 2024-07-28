package server

import "net"

// InitServer initializes a new server with the given listening address and maximum connections.
func InitServer(listenAddress string, maxConnections int) *Server {
	return &Server{
		listenAddress:   listenAddress, // Address on which the server listens
		ch:              make(chan struct{}), // General-purpose signal channel
		msgch:           make(chan Message, 10),  // Channel for broadcasting messages, buffered with a size of 10
		clients:         make(map[net.Conn]string),  // Map of active clients, with connection as key and username as value
		history:         make([]Message, 0),   // Slice to store chat history
		activeConnCount: make(chan struct{}, maxConnections), // Channel to limit the number of active connections
	}
}