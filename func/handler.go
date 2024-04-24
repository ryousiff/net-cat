package netcat

import (
	"bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func welcome() string {
	message, err := ioutil.ReadFile("penguin.txt")
	if err!= nil {
        log.Fatal("error displaying welcome image: ", err)
    }
	log.Printf("Welcome to TCP-Chat!\n")
	return string(message) + "\n"
}



func Handler (network net.Conn) {
	defer network.Close()

	msg := welcome()
	 
    network.Write([]byte(msg + "\n"))
	network.Write([]byte("Enter your name: "))
	Reader := bufio.NewReader(network)
	deadline := time.Now().Add(30 * time.Second)
	network.SetReadDeadline(deadline)



	Clients := &client {
		Reader: bufio.NewReader(network),
        Network: network,
	}
	Clients.Network.Write([]byte("Maximum clients reached."))

	Clients.Network.Close()

	_, err := Reader.ReadString('\n')
	if err!= nil {
        fmt.Printf("Connection failed: %v", err)
        return
    }

	if Reader == "" {
		network.Write([]byte("Please enter a name: "))
        _, err := Reader.ReadString('\n')
        if err!= nil {
            fmt.Printf("Connection failed: %v", err)
            return
        }
	}


}