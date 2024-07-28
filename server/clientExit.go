package server

import (
	"fmt"
	"net"
	"time"
)

func (s *Server) ClientExit(conn net.Conn) {
	s.mu.Lock()
	name := s.clients[conn]
	delete(s.clients, conn)
	s.mu.Unlock()
	s.msgch <- Message{msg: fmt.Sprintf("%s has left our chat...", name), sender: "System", time: time.Now()}
	conn.Close()
	<-s.activeConnCount // Release the connection slot
}


