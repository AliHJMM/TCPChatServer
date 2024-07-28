package server

import "fmt"

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error in connection: ", err)
			continue
		}

		// Check if max connections are reached
		select {
		case s.activeConnCount <- struct{}{}:
			// Allow the connection
			go s.ReadLoop(conn)
		default:
			// Reject the connection
			conn.Write([]byte("Server is full. Please try again later.\n"))
			conn.Close()
		}
	}
}
