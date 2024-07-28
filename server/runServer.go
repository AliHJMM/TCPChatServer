package server

import "net"


func (s *Server) RunServer() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	defer ln.Close()
	s.ln = ln
	go s.AcceptLoop()
	<-s.ch
	close(s.msgch)
	return nil
}