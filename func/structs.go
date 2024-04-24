package netcat

import (
    "net"
	"bufio"
)


type client struct {
	Network net.Conn 
	Name string
	Reader *bufio.Reader
}


var Clients client 