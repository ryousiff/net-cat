package netcat

import (
	"net"
	"sync"
)

type client struct {
	Network net.Conn
	Name    string
	// Reader *bufio.Reader
}

var (
	muclient   sync.Mutex
	allClients = make(map[string]*client)
	muhistory sync.Mutex
	history []string
)
var Clients client
var arrayofclient []client
