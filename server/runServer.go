package server

import "net"


// RunServer starts the server and listens for incoming TCP connections.
func (s *Server) RunServer() error {
	// Start listening on the specified address and port
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err // Return error if the server fails to start
	}

	defer ln.Close() // Ensure the listener is closed when the function exits
	s.ln = ln
	go s.AcceptLoop() // Start accepting connections in a separate goroutine
	<-s.ch // Wait for a signal on the channel (e.g., for shutdown)
	close(s.msgch) // Close the message channel after shutting down
	return nil
}