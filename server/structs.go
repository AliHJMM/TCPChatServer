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


