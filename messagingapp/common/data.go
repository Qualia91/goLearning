package common

import (
	"net"
	"sync"
)

type Message struct {
	Username string
	Message  string
	TimeSent string
}

type Database struct {
	Mut      sync.RWMutex
	Conns    []net.Conn
	MsgChan  chan string
	Messages []Message
}
