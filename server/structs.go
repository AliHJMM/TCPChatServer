package server

import (
	"net"
	"sync"
	"time"
)

// Message represents a single chat message.
type Message struct {
	msg    string // The content of the message
	sender string // The sender's name
	time   time.Time // The time the message was sent
}

// Server represents a TCP chat server.
type Server struct {
	listenAddress   string // Address and port the server listens on
	ln              net.Listener // TCP listener for incoming connections
	ch              chan struct{} // Channel for server control signals (e.g., shutdown)
	msgch           chan Message // Channel for broadcasting messages to clients
	clients         map[net.Conn]string // Map of active clients (connections and names)
	mu              sync.Mutex // Mutex to protect access to shared resources
	history         []Message // Slice to store the history of messages
	activeConnCount chan struct{} // Channel to track the number of active connections
}
