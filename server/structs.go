package server

import (
	"net"
	"sync"
	"time"
)

type Message struct {
	msg    string
	sender string
	time   time.Time
}

type Server struct {
	listenAddress   string
	ln              net.Listener
	ch              chan struct{}
	msgch           chan Message
	clients         map[net.Conn]string
	mu              sync.Mutex
	history         []Message
	activeConnCount chan struct{}
}
