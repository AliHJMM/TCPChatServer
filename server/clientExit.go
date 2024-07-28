package server

import (
	"fmt"
	"net"
	"time"
)

func (s *Server) ClientExit(conn net.Conn) {
	// Lock to safely access the clients map
	s.mu.Lock()
	name := s.clients[conn] // Get the client's name
	delete(s.clients, conn)  // Remove the client from the map
	s.mu.Unlock()

	// Notify all clients that this client has left
	s.msgch <- Message{msg: fmt.Sprintf("%s has left our chat...", name), sender: "System", time: time.Now()}
	conn.Close() // Close the connection
	<-s.activeConnCount// Decrement the active connection count
}


