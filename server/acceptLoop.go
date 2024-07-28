package server

import "fmt"

func (s *Server) AcceptLoop() {
	for {
		// Accept a new connection
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error in connection: ", err)
			continue
		}

		// Check if the maximum number of connections is reached
		select {
		case s.activeConnCount <- struct{}{}:
			// Allow the connection and start reading from it
			go s.ReadLoop(conn)
		default:
			// Reject the connection if the server is full
			conn.Write([]byte("Server is full. Please try again later.\n"))
			conn.Close()
		}
	}
}
