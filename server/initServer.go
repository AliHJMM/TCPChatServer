package server

import "net"

func InitServer(listenAddress string, maxConnections int) *Server {
	return &Server{
		listenAddress:   listenAddress,
		ch:              make(chan struct{}),
		msgch:           make(chan Message, 10),
		clients:         make(map[net.Conn]string),
		history:         make([]Message, 0),
		activeConnCount: make(chan struct{}, maxConnections),
	}
}